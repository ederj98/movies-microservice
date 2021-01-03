import * as React from 'react';
import { Route, Switch } from 'react-router-dom';
import { LazyFallback } from '../../shared/components/LazyFallback';

const MainPage = React.lazy(() => import('./pages/Main'));
const UpdateMainPage = React.lazy(() => import('./pages/MainActualizar'));

export const PeliculaRouter = () => (
  <React.Suspense fallback={<LazyFallback />}>
    {/* Layout compartido entre las rutas va aquí */}
    <Switch>
      <Route path="/" component={MainPage}></Route>
    </Switch>
  </React.Suspense>
);

export const PeliculaARouter = () => (
  <React.Suspense fallback={<LazyFallback />}>
    {/* Layout compartido entre las rutas va aquí */}
    <Switch>
      <Route path="/" component={UpdateMainPage}></Route>
    </Switch>
  </React.Suspense>
);
