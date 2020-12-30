import * as React from 'react';
import { Layout } from '../../../shared/components/Layout';
import { ProveedorGestionPeliculas } from '../hoc/ProveedorGestionPeliculas';
import { RouteComponentProps } from 'react-router-dom';

const MainPage: React.FC<RouteComponentProps> = () => {
  return (
    <Layout title="Peliculas" description="Gestión de peliculas">
      <ProveedorGestionPeliculas/>
    </Layout>
  );
};

MainPage.displayName = 'HomeMainPage';

export default MainPage;
