import logging

import time
from django_rq import job


@job
def register_user(form_data):
    logging.info('Received from data: %s' % (form_data,))
    time.sleep(5)
    # TODO: Register user
    # TODO: Send email
    logging.info('Tasks finished')
