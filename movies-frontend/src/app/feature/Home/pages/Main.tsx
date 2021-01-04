import * as React from 'react';
import { Layout } from 'app/shared/components/Layout';
import { RouteComponentProps } from 'react-router-dom';
import { ProveedorListarGestionPeliculas } from '../hoc/ProveedorListarPeliculas'

const HomeMainPage: React.FC<RouteComponentProps> = () => (
  <Layout title="Home" description="Home de la aplicación">
    <ProveedorListarGestionPeliculas/>
  </Layout>
);

HomeMainPage.displayName = 'HomeMainPage';

export default HomeMainPage;
