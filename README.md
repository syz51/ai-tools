# Translation Web

This is a web server written in Go using the Echo framework. It leverages htmx, tailwindcss, and templ for frontend templating.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to run the server:

- Go (version 1.21)
- node
- pnpm (optional)
- templ (go template library)
- air (optional for go hot reloading)

### Installing

A step by step series of examples that tell you how to get a development environment running:

1. Clone the repository
2. Navigate to the project directory
3. Run `pnpm install`
4. Run `pnpm dev`

Now you have the server running with tailwindcss watching css changes, and air for hot reloading

## Running the tests

```
go test ./...
```

## Deployment

Dockerfile is included so available for docker deployments

## Built With

- [Go](https://golang.org/) - The programming language used
- [Echo](https://echo.labstack.com/) - The web framework used

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

- **Roy Sun** - _Initial work_ - [syz51](https://github.com/syz51)

See also the list of [contributors](https://github.com/your-github-username/your-repo/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

- Hat tip to anyone whose code was used
- Inspiration
- etc
