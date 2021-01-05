import * as PropTypes from 'prop-types';
import * as React from 'react';
import { useEffect } from 'react';
import { DivContainer } from '../../../../shared/components/Div/index';
import { FormActualizarPelicula } from '../../components/ActualizarPelicula/index';
import { Pelicula } from '../../models/Pelicula';

interface ActualizarPeliculaProps {
  pelicula: Pelicula;
  buscarPelicula: (id: number) => void;
  actualizarPelicula: (pelicula: Pelicula) => void;
  id: number,
}

export const ActualizarPelicula: React.FC<ActualizarPeliculaProps> = ({
  pelicula,
  buscarPelicula,
  actualizarPelicula,
  id,
}) => {
  useEffect(() => {
    buscarPelicula(id);
  }, [id, buscarPelicula]);
  return (
    <DivContainer>
        <FormActualizarPelicula
          pelicula={pelicula}
          onSubmit={actualizarPelicula}
          formTitle="Actualizar pelicula"
        />
    </DivContainer>
  );
};

ActualizarPelicula.propTypes = {
  buscarPelicula: PropTypes.func.isRequired,
  actualizarPelicula: PropTypes.func.isRequired,
};
