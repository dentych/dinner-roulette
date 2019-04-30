# frontend

## Prerequisites

1. Node is installed on the system
    1. Download from NodeJS.org: https://nodejs.org/en/download/
1. Packages have been downloaded from NPM
    1. You can do this with `npm install` after Node has been installed
1. Backend is running
    1. See backend repo for further instruction on how to do this.
    
## Running the frontend

To run the frontend in development mode, with hot reload (reloads the page when changes have been made, without having
to restart the application), run:

```
npm run serve
```

## Building the frontend for production

Building for production should be done with Docker (docker build). However, it is possible to do it directly using
Node and NPM, which will generate the static HTML/CSS/JS files bundles needed.

```
npm run build
```

This defaults to `production` mode, which will enable all production features and variables.
It is possible to build for development or test using the parameter `--mode <mode>`

```
npm run build --mode development
```

Current modes available: production, test, development