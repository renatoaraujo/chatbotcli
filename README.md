# OpenAI GPT-3.5 Chatbot CLI

This is a command-line interface that allows you to interact with [OpenAI GPT-3.5](https://openai.com/gpt-3/) to get answers to your questions.

## Getting Started

Before using the chatbot, you will need to generate an API token from OpenAI. Follow these steps to do so:

1. Go to [OpenAI's website](https://openai.com/)
2. Click on **Sign Up** in the top-right corner of the page and create an account if you haven't already
3. Once logged in, go to the [API Tokens](https://beta.openai.com/dashboard/api-keys) page
4. Click on **New API Key** and give it a name
5. Copy the generated key

## Installation

To install the chatbot, run the following command:

```
make install token=<INSERT_YOUR_TOKEN_HERE>
```

This will build the chatbot and install your API token. Note that you need to replace `<INSERT_YOUR_TOKEN_HERE>` with the key you generated previously.

## Usage

To ask the chatbot a question, run the following command:

```
make ask
```

The chatbot will prompt you to enter your question. Type it in and hit Enter. The chatbot will then use OpenAI GPT-3.5 to generate an answer and display it on the screen.

## Contributing

If you would like to contribute to this project, please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
