
from rest_framework import routers
from django.conf.urls import url
from django.urls import path
from db.api.views import CatsView

app_name = 'myapi'

urlpatterns = [
	path('dogs/', CatsView.as_view(), name=""),
]

urlpatterns += router.urls
router = routers.SimpleRouter()
				