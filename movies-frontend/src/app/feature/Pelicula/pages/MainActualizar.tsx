import * as React from 'react';
import { Layout } from '../../../shared/components/Layout';
import { ProveedorActualizarPelicula } from '../hoc/ProveedorActualizarPelicula';
import { RouteComponentProps, useParams } from 'react-router-dom';

const UpdateMainPage: React.FC<RouteComponentProps> = () => {
  let {id} = useParams();
  console.log(id);
  return (
    <Layout title="Peliculas" description="Gestión de peliculas">
      <ProveedorActualizarPelicula id={2} />
    </Layout>
  );
};

UpdateMainPage.displayName = 'UpdateMainPage';

export default UpdateMainPage;
