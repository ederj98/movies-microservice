import * as PropTypes from 'prop-types';
import * as React from 'react';
import { DivContainer } from '../../../../shared/components/Div/index';
import { FormCrearPelicula } from '../../components/FormCrearPelicula';
import { Pelicula } from '../../models/Pelicula';

interface GestionPeliculasProps {
  agregarNuevaPelicula: (pelicula: Pelicula) => void;
}

export const GestionPeliculas: React.FC<GestionPeliculasProps> = ({
  agregarNuevaPelicula,
}) => {
  return (
    <DivContainer>
        <FormCrearPelicula
          onSubmit={agregarNuevaPelicula}
          formTitle="Crear pelicula"
        />
    </DivContainer>
  );
};

GestionPeliculas.propTypes = {
  agregarNuevaPelicula: PropTypes.func.isRequired,
};
