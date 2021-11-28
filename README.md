# CS5351FinalProj
Final Project of CS5351 Software Engineering\
Team 9\
Members: Gao Tao 55884812\
Liu Bingyu 56879209\
Ma Yixin 56885228\
Wang Jihua 57016790\
Li Zemin 56841822\
Li Bowen 56849383\
Wang Zhihao 56961113

# Client Implementation
The front-end of DEVils-Scrum-Tool is created with Vue.js, so it requires an installation of nodejs.
The project was developed using Vue CLI (@vue/cli). Open a new bash instance to run client:


## front-end plugins
    install Bootstrap,Element-ui,dataV,echarts,ViewUI
    
## Project setup
    npm install

## Compiles and hot-reloads for development
    npm run serve

## Up and Running!
Your client runs on http://localhost:9000/

## More configuration (optional)

### Compiles and minifies for production
    npm run build

### Run the unit tests (Jest)
    npm run test:unit

### Lints (code format) and fixes files
    npm run lint


## Other Deployment options

### Docker
    docker build -t client_vuejsapp .
    docker run -it -p 9000:80 --rm client_vuejsapp

### Heroku (for production)
1. Remove `/dist` directory from .gitignore
2. On `./src/main.js` apply `Vue.prototype.$url` as `https://your-backend.herokuapp.com/api`
3. Login to heroku and create a new project
4. Build locally, running `npm run build`
5. Git add, commit and push to heroku master



