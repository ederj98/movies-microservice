import * as PropTypes from 'prop-types';
import * as React from 'react';
import { DivContainer } from '../../../../shared/components/Div/index';
import { FormActualizarPelicula } from '../../components/ActualizarPelicula/index';
import { Pelicula } from '../../models/Pelicula';

interface ActualizarPeliculaProps {
  id: number,
  actualizarPelicula: (pelicula: Pelicula) => void;
}

export const ActualizarPelicula: React.FC<ActualizarPeliculaProps> = ({
  id,
  actualizarPelicula,
}) => {
  console.log(id)
  return (
    <DivContainer>
        <FormActualizarPelicula
          id={id}
          onSubmit={actualizarPelicula}
          formTitle="Actualizar pelicula"
        />
    </DivContainer>
  );
};

ActualizarPelicula.propTypes = {
  actualizarPelicula: PropTypes.func.isRequired,
};
