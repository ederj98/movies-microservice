import * as React from 'react';
import { Layout } from '../../../shared/components/Layout';
import { ProveedorCrearPeliculas } from '../hoc/ProveedorCrearPeliculas';
import { RouteComponentProps } from 'react-router-dom';

const MainPage: React.FC<RouteComponentProps> = () => {
  return (
    <Layout title="Peliculas" description="Gestión de peliculas">
      <ProveedorCrearPeliculas/>
    </Layout>
  );
};

MainPage.displayName = 'HomeMainPage';

export default MainPage;
