import * as React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import { HomeRouter } from 'app/feature/Home/HomeRouter';
import { NavigationHeader } from 'app/shared/components/NavigationHeader';
import { PeliculaRouter, PeliculaARouter } from 'app/feature/Pelicula/PeliculaRouter';

export const AppRouter = () => {
  return (
    <BrowserRouter>
      <NavigationHeader />
      <Switch>
        <Route path="/" exact component={HomeRouter} />
        <Route path="/home" component={HomeRouter} />
        <Route path="/movie" component={PeliculaRouter} />
        <Route path="/movies" component={PeliculaARouter} />
      </Switch>
    </BrowserRouter>
  );
};
