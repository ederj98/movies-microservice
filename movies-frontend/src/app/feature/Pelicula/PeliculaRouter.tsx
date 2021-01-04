import * as React from 'react';
import { Route, Switch, useRouteMatch } from 'react-router-dom';
import { LazyFallback } from '../../shared/components/LazyFallback';

const MainPage = React.lazy(() => import('./pages/Crear'));
const UpdateMainPage = React.lazy(() => import('./pages/Actualizar'));

export const PeliculaRouter = () => (
  <React.Suspense fallback={<LazyFallback />}>
    {/* Layout compartido entre las rutas va aquí */}
    <Switch>
      <Route path="/" component={MainPage}></Route>
    </Switch>
  </React.Suspense>
);

export const PeliculaARouter = () => {
  const { path } = useRouteMatch();
  return (
    <React.Suspense fallback={<LazyFallback />}>
    {/* Layout compartido entre las rutas va aquí */}
    <Switch>
      <Route path={`${path}/:id`} component={UpdateMainPage}></Route>
    </Switch>
  </React.Suspense>
  )
  };
