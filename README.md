# Startup

Starting the two servers should be done using two terminal windows, both starting in the directory of the installed project.
Start by starting the **back-end** server **first**, when the backend is started up last, just reload the front-end when the back-end has started.

## Back-end

For the backend, you need to start up the server by doing:
```bash
cd api
go run *.go
```

## front-end

To start up the front end do the following:
```bash
cd client
npm install
```

This will install the dependencies needed to start or build the front-end server.

To run this server do either: `npm run dev` for the development server or do:
```bash
npm run build
npm run start
```


# Usage

If both servers are started correctly, you can go to `localhost:5000`, where you can find a page with an input field, and a list of movies already in the database.

## Display a movie bigger (with poster)
To see a movie displayed bigger or you would like to see the poster, just click one of the movies.

## Finding another movie
To see a new movie, just input an IMDB id into the search bar and press enter or the *search* button
