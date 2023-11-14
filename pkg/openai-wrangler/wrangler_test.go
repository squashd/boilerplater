package wrangler

import (
	"fmt"
	"testing"

	"github.com/SQUASHD/boilerplater/internal/shared/models"
)

func TestFindCodeBlockIndices(t *testing.T) {
	type args struct {
		content            string
		start              string
		end                string
		expectedStartIndex int
		expectedEndIndex   int
		expectedErr        bool
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Empty code block",
			args: args{
				content:            "```````",
				start:              "```",
				end:                "```",
				expectedStartIndex: 3,
				expectedEndIndex:   3,
				expectedErr:        false,
			},
		},
		{
			name: "Short code block",
			args: args{
				content:            "````",
				start:              "```",
				end:                "```",
				expectedStartIndex: -1,
				expectedEndIndex:   -1,
				expectedErr:        true,
			},
		},
		{
			name: "Standard JSON block",
			args: args{
				content:            "```json\n{\"test\": \"test\"}\n```",
				start:              "```",
				end:                "```",
				expectedStartIndex: 3,
				expectedEndIndex:   25,
				expectedErr:        false,
			},
		},
		{
			name: "Preambled JSON block",
			args: args{
				content:            "Preambling```json\n{\"test\": \"test\"}\n```",
				start:              "```json",
				end:                "```",
				expectedStartIndex: 17,
				expectedEndIndex:   35,
				expectedErr:        false,
			},
		},
		{
			name: "Mismatched start and end tags",
			args: args{
				content:            "```json\n{\"test\": \"test\"}\n```",
				start:              "```",
				end:                "```json",
				expectedStartIndex: -1,
				expectedEndIndex:   -1,
				expectedErr:        true,
			},
		},
		{
			name: "Code block with no end tag",
			args: args{
				content:            "```json\n{\"test\": \"test\"}",
				start:              "```",
				end:                "```",
				expectedStartIndex: -1,
				expectedEndIndex:   -1,
				expectedErr:        true,
			},
		},
		{
			name: "Code block with no start tag",
			args: args{
				content:            "{\"test\": \"test\"}\n```",
				start:              "```",
				end:                "```",
				expectedStartIndex: -1,
				expectedEndIndex:   -1,
				expectedErr:        true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start, stop, err := findCodeBlockIndices(tt.args.content, tt.args.start, tt.args.end)
			if start != tt.args.expectedStartIndex {
				t.Errorf("Test %v failed: Expected start index to be %d, got %d", tt.name, tt.args.expectedStartIndex, start)
			}
			if stop != tt.args.expectedEndIndex {
				t.Errorf("Test %v failed: Expected end index to be %d, got %d", tt.name, tt.args.expectedEndIndex, stop)
			}
			if (err != nil) != tt.args.expectedErr {
				t.Errorf("Test %v failed: Expected error to be %v, got %v", tt.name, tt.args.expectedErr, err)
			}
		})
	}
}

func TestExtractAndCleanJSON(t *testing.T) {
	type args struct {
		content        string
		expectedResult string
		expectedErr    bool
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Empty JSON",
			args: args{
				content:        "```json\\n```",
				expectedResult: "",
				expectedErr:    false,
			},
		},
		{
			name: "Standard JSON",
			args: args{
				content:        "```json\\n{\"test\": \"test\"}\\n```",
				expectedResult: "{\"test\": \"test\"}",
				expectedErr:    false,
			},
		},
		{
			name: "Preambled JSON",
			args: args{
				content:        "Preambling```json\\n{\"test\": \"test\"}\\n```",
				expectedResult: "{\"test\": \"test\"}",
				expectedErr:    false,
			},
		},
		{
			name: "JSON with newlines",
			args: args{
				content:        "```json\\n{\"test\": \"test\",\\n\"test2\": \"test2\"}\\n```",
				expectedResult: "{\"test\": \"test\",\"test2\": \"test2\"}",
				expectedErr:    false,
			},
		},
		{
			name: "Poorly formatted JSON",
			args: args{
				content:        "```json\\n{\"test\": \"test\",\\n\"test2\": \"test2\"}\\n```",
				expectedResult: "{\"test\": \"test\",\"test2\": \"test2\"}",
				expectedErr:    false,
			},
		},
		{
			name: "JSON with no end tag",
			args: args{
				content:        "```json\\n{\"test\": \"test\",\\n\"test2\": \"test2\"}",
				expectedResult: "",
				expectedErr:    true,
			},
		},
		{
			name: "JSON with no start tag",
			args: args{
				content:        "{\"test\": \"test\",\\n\"test2\": \"test2\"}\\n```",
				expectedResult: "",
				expectedErr:    true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result, err := extractAndCleanJSON(tt.args.content)
			if result != tt.args.expectedResult {
				t.Errorf("Test %v failed: Expected result to be %v, got %v", tt.name, tt.args.expectedResult, result)
			}
			if (err != nil) != tt.args.expectedErr {
				t.Errorf("Test %v failed: Expected error to be %v, got %v", tt.name, tt.args.expectedErr, err)
			}

		})
	}
}

func TestProcessContent(t *testing.T) {
	type args struct {
		content     string
		model       interface{}
		expectedErr bool
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Correctly Formatted JSON",
			args: args{
				content:     "Sure, here's a possible outline for your project using the JSON structure you provided:\n\n```json\n{\n  \"Title\": \"Time Keeper Web Service\",\n  \"Objective\": \"Develop a web service using Go that allows users to create, read, and manage timekeeping records. This project should be of production quality, adhering to best practices in coding, testing, and version control.\",\n  \"Features\": [\n    {\n      \"Name\": \"User Authentication\",\n      \"Description\": \"Users should be able to securely log in to the web service\",\n      \"Tips\": [\n        \"Consider using an external library for handling authentication\",\n        \"Ensure passwords are stored securely\"\n      ]\n    },\n    {\n      \"Name\": \"Time Record Creation\",\n      \"Description\": \"Users should be able to create new time records\",\n      \"Tips\": [\"Consider the different pieces of information you will need to store for each time record\"]\n    },\n    {\n      \"Name\": \"Time Record Management\",\n      \"Description\": \"Users should be able to view, update, and delete their existing time records\",\n      \"Tips\": [\"Consider how you will handle changes to time records\"]\n    },\n    {\n      \"Name\": \"Data Persistence\",\n      \"Description\": \"Time records should be stored persistently, so they persist across sessions\",\n      \"Tips\": [\n        \"Consider using a database for this feature\",\n        \"Ensure that the database is secured and that only authenticated users can make changes to their own time records\"\n      ]\n    }\n  ],\n  \"Steps\": [\n    {\n      \"Description\": \"Set up your development environment\",\n      \"Tips\": \"Install Go and any necessary libraries\"\n    },\n    {\n      \"Description\": \"Design the data structure for your time records\",\n      \"Tips\": \"Think about what information you need to store for each time record\"\n    },\n    {\n      \"Description\": \"Implement user authentication\",\n      \"Tips\": \"Consider using an external library to handle this\"\n    },\n    {\n      \"Description\": \"Implement the ability to create, read, update, and delete time records\",\n      \"Tips\": \"Remember to validate user input\"\n    },\n    {\n      \"Description\": \"Implement data persistence\",\n      \"Tips\": \"Consider using a database for this\"\n    },\n    {\n      \"Description\": \"Test your web service\",\n      \"Tips\": \"Make sure to test all features of your web service, including edge cases\"\n    }\n  ],\n  \"Setup\": \"Install Go and any necessary libraries. Set up a database for storing time records.\",\n  \"Testing\": \"Write unit tests for each function in your web service. Consider using a tool like Postman for testing your API.\",\n  \"Debugging\": \"Use Go's built-in debugging tools to help diagnose and fix issues\",\n  \"Extras\": [\"Consider adding an API for your web service\", \"Consider adding a front-end for your web service\"]\n}\n```\n\nThis outline provides a roadmap for your project, starting with setting up your development environment, designing your data structure, and implementing each feature one by one. It suggests using external libraries for user authentication and potentially for data persistence as well. It also emphasizes the importance of testing your web service and includes a step for this. Finally, it provides some extra ideas for enhancing your web service, such as adding an API or a front-end.",
				model:       models.IntermediateProject{},
				expectedErr: false,
			},
		},
		{
			name: "Incorrectly Formatted JSON",
			args: args{
				content:     "",
				model:       models.IntermediateProject{},
				expectedErr: true,
			},
		},
		{
			name: "JSON with no code block tags",
			args: args{
				content:     "{\n  \"title\": \"Spotify Playlist Creator\",\n  \"objective\": \"Learn how to interact with APIs, handle data, and utilize basic JavaScript by creating a Spotify playlist creator.\",\n  \"steps\": [\n    {\n      \"description\": \"Get to know the Spotify API\",\n      \"tips\": \"Start by reading the Spotify API documentation. Understand how authentication works and look into the endpoints that allow you to search for tracks and create playlists.\"\n    },\n    {\n      \"description\": \"Set up your project\",\n      \"tips\": \"Create a new directory for your project. Initialize it as a Node.js project and install necessary dependencies like 'express' for setting up the server and 'axios' for making HTTP requests.\"\n    },\n    {\n      \"description\": \"Implement Spotify Authentication\",\n      \"tips\": \"Use the 'Client Credentials' flow as it's suitable for beginners. You'll need to sign up as a Spotify Developer to get your client ID and secret.\"\n    },\n    {\n      \"description\": \"Create routes to search for tracks\",\n      \"tips\": \"Use the '/v1/search' endpoint of Spotify API. Make sure to handle errors and edge cases like no results found.\"\n    },\n    {\n      \"description\": \"Create routes to create a playlist\",\n      \"tips\": \"Use the '/v1/users/{user_id}/playlists' endpoint. You might need to handle permissions to modify the user's playlists.\"\n    },\n    {\n      \"description\": \"Test your application\",\n      \"tips\": \"Use Postman or any other API testing tool to make sure your routes are working as expected.\"\n    }\n  ],\n  \"watchOuts\": [\n    \"Spotify API rate limits. Make sure not to exceed the allowed number of requests per minute.\",\n    \"Handle errors properly. If a request fails, your application should be able to recover from it.\",\n    \"Keep your Spotify client ID and secret safe. Do not push them to public repositories.\"\n  ],\n  \"extraChallenges\": [\n    \"Add a feature to remove tracks from a playlist.\",\n    \"Allow users to log in with their Spotify account and modify their own playlists.\",\n    \"Implement a front-end with a user-friendly interface.\"\n  ]\n}",
				model:       models.BeginnerProject{},
				expectedErr: false,
			},
		},
		{
			name: "Beginner Project",
			args: args{
				content:     "{\n  \"title\": \"Weather App using OpenWeatherMap API in Python\",\n  \"objective\": \"Build a simple weather forecast application using the OpenWeatherMap API in Python. The application will allow users to get the current weather data for any city in the world.\",\n  \"steps\": [\n    {\n      \"description\": \"Setup your Python development environment. If you haven't already, install Python and a code editor of your choice.\",\n      \"tips\": \"Python can be downloaded from the official website. Some popular choices for a code editor are PyCharm, VS Code or Atom.\"\n    },\n    {\n      \"description\": \"Sign up for a free OpenWeatherMap account to get your API key.\",\n      \"tips\": \"The API key will be used to authenticate your requests. Be careful not to share this key publicly.\"\n    },\n    {\n      \"description\": \"Install the necessary Python libraries. For this project, you'll need requests and json.\",\n      \"tips\": \"These libraries can be installed using pip - Python's package installer. The command is `pip install requests json`.\"\n    },\n    {\n      \"description\": \"Create a new Python file and import the libraries you just installed.\",\n      \"tips\": \"Use the import statement at the top of your Python file to do this. For example: `import requests, json`.\"\n    },\n    {\n      \"description\": \"Write a function to make a GET request to the OpenWeatherMap API. The function should take in a city name and return the weather data for that city.\",\n      \"tips\": \"Use the requests library to make the API call. The API endpoint URL should look something like this: `http://api.openweathermap.org/data/2.5/weather?q={city name}&appid={your api key}`.\"\n    },\n    {\n      \"description\": \"In the same function, parse the JSON response to get the information you want to display. This could include the city name, temperature, weather description, etc.\",\n      \"tips\": \"Use the json library to parse the response. You can access individual pieces of data using the keys in the JSON object. For example, `json_data['name']` would get the city name.\"\n    },\n    {\n      \"description\": \"Finally, create a user interface in the console. Prompt the user to enter a city name, then call your function to get the weather data and display it.\",\n      \"tips\": \"Use the input function to get user input. Make sure to handle any errors that might occur, such as the user entering a city name that doesn't exist.\"\n    }\n  ],\n  \"watchOuts\": [\n    \"Make sure to handle any errors that might occur, such as the user entering a city name that doesn't exist or the API being unavailable.\",\n    \"Be careful not to share your OpenWeatherMap API key publicly.\",\n    \"Remember that the weather data is in Kelvin by default. You might want to convert it to Celsius or Fahrenheit before displaying it.\"\n  ],\n  \"extraChallenges\": [\n    \"Extend the app to display the weather forecast for the next few days.\",\n    \"Add the ability for users to save their favorite cities and quickly check the weather in those places.\",\n    \"Create a graphical user interface using a library like Tkinter.\"\n  ]\n}",
				model:       models.BeginnerProject{},
				expectedErr: false,
			},
		},
		{
			name: "Any project, any language",
			args: args{
				content:     "Sure, here's a project outline for a simple \"Weather Forecast App\" using Python and a weather API. Python is a great language for beginners due to its readability and straightforward syntax.\n\n```json\n{\n  \"Title\": \"Weather Forecast App\",\n  \"Objective\": \"Create a simple console-based weather forecast application using Python and OpenWeatherMap's API\",\n  \"Steps\": [\n    {\n      \"Description\": \"Setup your Python environment\",\n      \"Tips\": \"If you haven't already, install Python and a code editor (recommend VS Code or PyCharm). You can find numerous tutorials online on how to set these up.\"\n    },\n    {\n      \"Description\": \"Learn the basics of APIs\",\n      \"Tips\": \"An API (Application Programming Interface) is a way for different software applications to communicate with each other. In this project, we will be using the OpenWeatherMap's API to get weather data.\"\n    },\n    {\n      \"Description\": \"Register for a free API key from OpenWeatherMap\",\n      \"Tips\": \"Go to the OpenWeatherMap's website, create an account, and generate a free API key. Remember to save this key as you will need it to make requests.\"\n    },\n    {\n      \"Description\": \"Start your Python project\",\n      \"Tips\": \"Create a new Python file (e.g., weather_app.py) in your code editor.\"\n    },\n    {\n      \"Description\": \"Install necessary Python libraries\",\n      \"Tips\": \"We will need the 'requests' library to make HTTP requests. You can install it using pip - 'pip install requests' in the terminal.\"\n    },\n    {\n      \"Description\": \"Write a function to get weather data\",\n      \"Tips\": \"Use the 'requests' library to send a GET request to OpenWeatherMap's API with your API key and the city name as parameters. The API will return a JSON response with the weather data.\"\n    },\n    {\n      \"Description\": \"Write a function to display weather data\",\n      \"Tips\": \"This function will take the JSON response from the previous step and print the relevant weather information in a user-friendly format.\"\n    },\n    {\n      \"Description\": \"Add user input functionality\",\n      \"Tips\": \"Allow users to input the city name for which they want to see the weather forecast.\"\n    }\n  ],\n  \"WatchOuts\": [\n    \"Make sure to handle exceptions and errors properly, for example, when the city name entered by the user is not found.\",\n    \"Be careful with your API key - do not share it publicly or push it to GitHub.\",\n    \"The OpenWeatherMap's free API key has limitations, such as a limited number of requests per minute.\"\n  ],\n  \"ExtraChallenges\": [\n    \"Try to add more features to the app, like showing the forecast for the next few days instead of just the current weather.\",\n    \"Refactor your code into classes to practice Object-Oriented Programming (OOP).\",\n    \"Build a GUI (Graphical User Interface) for the app using a library like Tkinter.\"\n  ]\n}\n```\n\nThis project will help you solidify your understanding of basic programming concepts, such as functions, API usage, error handling, and user input.",
				model:       models.BeginnerProject{},
				expectedErr: false,
			},
		},
		{
			name: "Currency converter in Rust",
			args: args{
				content:     "{\n  \"title\": \"Currency Converter Service\",\n  \"objective\": \"Build a simple currency converter service using Rust and the ExchangeRate-API.\",\n  \"steps\": [\n    {\n      \"description\": \"Set up your Rust programming environment\",\n      \"tips\": \"Make sure you have Rust installed on your system. You can download it from the official website. You will also need the 'cargo' tool for building and managing Rust projects.\"\n    },\n    {\n      \"description\": \"Learn about API interaction in Rust\",\n      \"tips\": \"You'll be using the 'reqwest' library to send HTTP requests. Understanding how to send GET requests and handle responses will be crucial for interacting with the ExchangeRate-API.\"\n    },\n    {\n      \"description\": \"Register for the ExchangeRate-API\",\n      \"tips\": \"Go to the ExchangeRate-API website and create a free account. After you've registered, you'll be given an API key. This key will be used in your requests to access currency data.\"\n    },\n    {\n      \"description\": \"Create a new Rust project\",\n      \"tips\": \"Use the 'cargo new' command to create a new project. The command will generate a new directory with the necessary files and structure for a Rust project.\"\n    },\n    {\n      \"description\": \"Implement the API request\",\n      \"tips\": \"Use the 'reqwest' library to send a GET request to the ExchangeRate-API. Remember to include your API key in the request.\"\n    },\n    {\n      \"description\": \"Parse the API response\",\n      \"tips\": \"The 'serde' library will help you to parse the JSON response from the API. You'll need to create a data structure that matches the response format.\"\n    },\n    {\n      \"description\": \"Implement the currency conversion\",\n      \"tips\": \"Use the data from the API response to perform a currency conversion. Make sure to handle potential errors, like if the user inputs an invalid currency code.\"\n    },\n    {\n      \"description\": \"Test your program\",\n      \"tips\": \"Try out different currency conversions to make sure your program is working correctly. Pay attention to how your program handles invalid input and edge cases.\"\n    }\n  ],\n  \"watchOuts\": [\n    \"Ensure you handle potential errors during the API request, such as network issues or an invalid API key.\",\n    \"Make sure your program can handle invalid currency codes, either by checking the input against a list of valid codes or by handling the error returned by the API.\"\n  ],\n  \"extraChallenges\": [\n    \"Add a user interface, either command-line or graphical, to your program.\",\n    \"Allow the user to convert between multiple different currencies at once.\",\n    \"Implement a 'history' feature that keeps track of the user's past conversions.\"\n  ]\n}",
				model:       models.BeginnerProject{},
				expectedErr: false,
			},
		},
		{
			name: "Extra field causing err",
			args: args{
				content:     "Sure, here is a project outline that fits your requirements:\n\n```json\n{\n  \"IntermediateProject\": {\n    \"Title\": \"Time Keeping Web Service\",\n    \"Objective\": \"To create a web service in Go that handles start/stop time tracking requests and sends records to an employer.\",\n    \"Features\": [\n      {\n        \"Name\": \"API Endpoints\",\n        \"Description\": \"Create RESTful API endpoints for start/stop time tracking.\",\n        \"Tips\": [\"Use a Go web framework like Gin or Echo to handle routing and middleware.\", \"Design your endpoints around resources, following REST principles.\"]\n      },\n      {\n        \"Name\": \"Data Persistence\",\n        \"Description\": \"Implement a database to store time tracking records.\",\n        \"Tips\": [\"Consider using an ORM library for Go like GORM to interact with your database.\", \"SQLite could be a good starting point for your database.\"]\n      },\n      {\n        \"Name\": \"Authentication\",\n        \"Description\": \"Secure your API with authentication.\",\n        \"Tips\": [\"JWT authentication is a common method for securing APIs.\", \"Consider using a Go JWT library like jwt-go.\"]\n      },\n      {\n        \"Name\": \"Email Sending\",\n        \"Description\": \"Implement a feature that sends time tracking records to your employer via email.\",\n        \"Tips\": [\"Use a Go SMTP library to send emails.\", \"Remember to configure your email server securely.\"]\n      }\n    ],\n    \"Steps\": [\n      {\n        \"Description\": \"Set up your development environment for Go.\",\n        \"Tips\": \"You need the Go compiler and a text editor or IDE that supports Go. Visual Studio Code with the Go extension is a good choice.\"\n      },\n      {\n        \"Description\": \"Create your API endpoints.\",\n        \"Tips\": \"Start with a single resource, like a time tracking record. Create endpoints for starting and stopping time tracking.\"\n      },\n      {\n        \"Description\": \"Implement data persistence.\",\n        \"Tips\": \"Design your database schema based on the resources your API handles. Use migrations to manage your database schema.\"\n      },\n      {\n        \"Description\": \"Add authentication to your API.\",\n        \"Tips\": \"Secure your endpoints so that only authenticated requests can start or stop time tracking.\"\n      },\n      {\n        \"Description\": \"Implement email sending.\",\n        \"Tips\": \"Test sending emails with a dummy account before configuring your real email server.\"\n      },\n      {\n        \"Description\": \"Deploy your web service.\",\n        \"Tips\": \"You can use a platform like Heroku or Google Cloud Run to deploy your web service. Don't forget to secure your deployed service with SSL.\"\n      }\n    ],\n    \"Setup\": \"Install the Go compiler, set up your text editor or IDE for Go development, and install any necessary Go libraries.\",\n    \"Testing\": \"Write unit tests for your API endpoints, data persistence logic, and email sending functionality. Use a Go testing framework like GoConvey.\",\n    \"Debugging\": \"Use a Go debugger like Delve to step through your code and find bugs. Make use of logging to trace how your code is executing.\",\n    \"Extras\": [\n      \"Add validation to your API to check for invalid or missing data in requests.\",\n      \"Implement rate limiting to protect your API from abuse.\",\n      \"Add an admin interface to manage time tracking records.\"\n    ]\n  }\n}\n```",
				model:       models.IntermediateProject{},
				expectedErr: true,
			},
		},
		{
			name: "Wrong model specified",
			args: args{
				content:     "Sure, here's a possible outline for your project using the JSON structure you provided:\n\n```json\n{\n  \"Title\": \"Time Keeper Web Service\",\n  \"Objective\": \"Develop a web service using Go that allows users to create, read, and manage timekeeping records. This project should be of production quality, adhering to best practices in coding, testing, and version control.\",\n  \"Features\": [\n    {\n      \"Name\": \"User Authentication\",\n      \"Description\": \"Users should be able to securely log in to the web service\",\n      \"Tips\": [\n        \"Consider using an external library for handling authentication\",\n        \"Ensure passwords are stored securely\"\n      ]\n    },\n    {\n      \"Name\": \"Time Record Creation\",\n      \"Description\": \"Users should be able to create new time records\",\n      \"Tips\": [\"Consider the different pieces of information you will need to store for each time record\"]\n    },\n    {\n      \"Name\": \"Time Record Management\",\n      \"Description\": \"Users should be able to view, update, and delete their existing time records\",\n      \"Tips\": [\"Consider how you will handle changes to time records\"]\n    },\n    {\n      \"Name\": \"Data Persistence\",\n      \"Description\": \"Time records should be stored persistently, so they persist across sessions\",\n      \"Tips\": [\n        \"Consider using a database for this feature\",\n        \"Ensure that the database is secured and that only authenticated users can make changes to their own time records\"\n      ]\n    }\n  ],\n  \"Steps\": [\n    {\n      \"Description\": \"Set up your development environment\",\n      \"Tips\": \"Install Go and any necessary libraries\"\n    },\n    {\n      \"Description\": \"Design the data structure for your time records\",\n      \"Tips\": \"Think about what information you need to store for each time record\"\n    },\n    {\n      \"Description\": \"Implement user authentication\",\n      \"Tips\": \"Consider using an external library to handle this\"\n    },\n    {\n      \"Description\": \"Implement the ability to create, read, update, and delete time records\",\n      \"Tips\": \"Remember to validate user input\"\n    },\n    {\n      \"Description\": \"Implement data persistence\",\n      \"Tips\": \"Consider using a database for this\"\n    },\n    {\n      \"Description\": \"Test your web service\",\n      \"Tips\": \"Make sure to test all features of your web service, including edge cases\"\n    }\n  ],\n  \"Setup\": \"Install Go and any necessary libraries. Set up a database for storing time records.\",\n  \"Testing\": \"Write unit tests for each function in your web service. Consider using a tool like Postman for testing your API.\",\n  \"Debugging\": \"Use Go's built-in debugging tools to help diagnose and fix issues\",\n  \"Extras\": [\"Consider adding an API for your web service\", \"Consider adding a front-end for your web service\"]\n}\n```\n\nThis outline provides a roadmap for your project, starting with setting up your development environment, designing your data structure, and implementing each feature one by one. It suggests using external libraries for user authentication and potentially for data persistence as well. It also emphasizes the importance of testing your web service and includes a step for this. Finally, it provides some extra ideas for enhancing your web service, such as adding an API or a front-end.",
				model:       models.BeginnerProject{},
				expectedErr: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var model interface{}

			switch tt.args.model.(type) {
			case models.IntermediateProject:
				model = &models.IntermediateProject{}
			case models.BeginnerProject:
				model = &models.BeginnerProject{}
			default:
				t.Fatalf("Unsupported model type")
			}

			err := processContent(tt.args.content, model)
			fmt.Printf("%+v\n", model)
			if (err != nil) != tt.args.expectedErr {
				t.Errorf("Test %v failed: Expected error to be %v, got %v", tt.name, tt.args.expectedErr, err)
			}

		})

	}
}
