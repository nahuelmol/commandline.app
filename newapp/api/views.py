
from rest_framework.response import Response
from rest_framework import viewsets, status
from rest_framework.decorators import action
from rest_framework import permissions, generics, authentication

from django.shortcuts import get_object_or_404
from django.contrib.auth.models import User

from .permissions import IsUserLoggedIn
from .models import Cat
from .api.serializer import CatSerializer

class CatsView(viewsets.ViewSet):
	def create(self, request):
		pass
				