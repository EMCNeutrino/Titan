import re

from django import template
from django.core.urlresolvers import reverse, NoReverseMatch

register = template.Library()


# @register.simple_tag
# def active(request, pattern):
#     try:
#         p = re.compile(pattern)
#         m = p.match(request.path)
#         if m:
#             return 'active'
#     except:
#         pass
#     return ''

@register.simple_tag(takes_context=True)
def active(context, pattern_or_urlname):
    try:
        pattern = '^' + reverse(pattern_or_urlname)
    except NoReverseMatch:
        pattern = pattern_or_urlname
    path = context['request'].path
    if re.search(pattern, path):
        return 'active'
    return ''