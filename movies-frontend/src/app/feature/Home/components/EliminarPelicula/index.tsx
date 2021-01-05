import * as PropTypes from 'prop-types';
import * as React from 'react';
import { Button } from '../../../../shared/components/Button';
import { Pelicula } from '../../../Pelicula/models/Pelicula';

interface BtnEliminarPeliculaProps {
  onEliminar: (pelicula: Pelicula) => any;
  pelicula: Pelicula;
}

export const BtnEliminarPelicula: React.FC<BtnEliminarPeliculaProps> = ({
  onEliminar,
  pelicula,
}) => {
  const handleEliminar = () => onEliminar(pelicula);
  return (
    <Button onClick={handleEliminar}>
      <span role="img" aria-labelledby="trash">
        🗑️
      </span>
    </Button>
  );
};

BtnEliminarPelicula.propTypes = {
  pelicula: PropTypes.shape({
    id: PropTypes.number.isRequired,
    name: PropTypes.string.isRequired,
    director: PropTypes.string.isRequired,
    writer: PropTypes.string.isRequired,
    stars: PropTypes.string.isRequired,
  }).isRequired,
  onEliminar: PropTypes.func.isRequired,
};
