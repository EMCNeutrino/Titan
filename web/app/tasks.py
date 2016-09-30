from __future__ import absolute_import

from celery import task
from celery.utils.log import get_task_logger

logger = get_task_logger(__name__)


@task
def register_user(data):
    logger.info('Data: %s' % (data,))
    return
