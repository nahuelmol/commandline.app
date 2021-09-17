package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"github.com/thatisuday/commando"
)

func main() {
	commando.
		SetExecutableName("cracky").
		SetVersion("1.0.0").
		SetDescription("This tool lists the contents of a directory in tree-like format.\nIt can also display information about files and folders like size, permission and ownership.")

	// configure the root command
	commando.
		Register(nil).
		AddArgument("dir", "local directory path", "./").                     // default `./`
		AddFlag("level,l", "level of depth to travel", commando.Int, 1).      // default `1`
		AddFlag("size", "display size of the each file", commando.Bool, nil). // default `false`
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("Printing options of the `dir` command...\n\n")

			// print arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

			// print flags
			for k, v := range flags {
				fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})

	// configure info command
	commando.
		Register("info").
		SetShortDescription("displays detailed information of a directory").
		SetDescription("This command displays more information about the contents of the directory like size, permission and ownership, etc.").
		AddArgument("dir", "local directory path", "./").                  // default `./`
		AddFlag("level,l", "level of depth to travel", commando.Int, nil). // required
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("Printing options of the `info` command...\n\n")

			// print arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

			// print flags
			for k, v := range flags {
				fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})

	commando.
		Register("builder").
		SetShortDescription("just create folders and files for an API creation in django").
		SetDescription("the same as the short one").
		AddArgument("app", "the app where is located the API", "api").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue){
			fmt.Printf("There will be the logic for directories creation...\n\n")


			if args["app"].Value == "api" {
				fmt.Println("you must provide the django app's name")
			}else {
				var end_dir string = args["app"].Value + "/api"
				err:= os.Mkdir(end_dir, 0755)
				if err != nil {
					log.Fatal(err)
				}

				var serializer_file string = end_dir + "/serializer.py" 
				var views_file		string = end_dir + "/views.py"
				var url_file		string = end_dir + "/url.py"

				var url_c		string = `
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
				`

				var serializer_c string = `
from rest_framework import serializers
from django.contrib.auth.models import User
from db.models import Cat

class CatSerializer(serializers.ModelSerializer):
	class Meta:
		fields	= '__all__'
		model 	= Cat 
				`

				var views_c string = `
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
				`

				var f_contents 		= [3]string {serializer_c, views_c, url_c}
				var files 			= [3]string {serializer_file, views_file, url_file}

				for content, file := range files {
					fmt.Println(file)

					err_ := ioutil.WriteFile(file, []byte(f_contents[content]), 0755)
   					if err_ != nil {
        				fmt.Printf("Unable to write file: %v", err)
    				}

				}

				
			}

		})

	commando.Parse(nil)
}