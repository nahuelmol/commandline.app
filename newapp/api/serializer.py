
from rest_framework import serializers
from django.contrib.auth.models import User
from db.models import Cat

class CatSerializer(serializers.ModelSerializer):
	class Meta:
		fields	= '__all__'
		model 	= Cat 
				