import json
from pathlib import Path

import azure.functions as func
import elasticapm
import requests
from elasticapm.contrib.serverless.azure import ElasticAPMExtension
from elasticapm.handlers.structlog import structlog_processor
from structlog import WriteLogger, wrap_logger
from structlog.processors import JSONRenderer

log_file = Path("traces.log")
wrapped_logger = WriteLogger(file=log_file.open("at"))
logger = wrap_logger(wrapped_logger, processors=[structlog_processor, JSONRenderer()])
log = logger.new()

ElasticAPMExtension.configure(
    server_url="http://localhost:8200",
    service_name="functions-app",
    environment="development",
    use_structlog=True,
)


app = func.FunctionApp()


@app.route(route="health")
def health(req: func.HttpRequest) -> func.HttpResponse:
    log.msg("health check started")
    response = requests.get("https://google.com")
    response_data = build_response(response.status_code == 200)

    elasticapm.label(my_status=response.status_code)
    log.msg("health check completed")
    return func.HttpResponse(json.dumps(response_data), status_code=200)


@elasticapm.capture_span()
def build_response(is_healthy: bool) -> dict:
    log.msg("build_response started")
    return {"status": "UP" if is_healthy else "DOWN"}
