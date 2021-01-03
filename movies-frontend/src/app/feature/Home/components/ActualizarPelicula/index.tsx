import * as PropTypes from 'prop-types';
import * as React from 'react';
import { useHistory } from 'react-router-dom';
import { Button } from '../../../../shared/components/Button';
import { Pelicula } from '../../../Pelicula/models/Pelicula';

interface BtnActualizarPeliculaProps {
  onActualizar: (pelicula: Pelicula) => any;
  pelicula: Pelicula;
}

export const BtnActualizarPelicula: React.FC<BtnActualizarPeliculaProps> = ({
  onActualizar,
  pelicula,
}) => {
  const history = useHistory();
  const handleActualizar = () => history.push(`/movies/${pelicula.Id}`);
  return (
    <Button onClick={handleActualizar}>
      <span role="img" aria-labelledby="trash">
      📝
      </span>
    </Button>
  );
};

BtnActualizarPelicula.propTypes = {
  pelicula: PropTypes.shape({
    Id: PropTypes.number.isRequired,
    Name: PropTypes.string.isRequired,
    Director: PropTypes.string.isRequired,
    Writer: PropTypes.string.isRequired,
    Stars: PropTypes.string.isRequired,
  }).isRequired,
  onActualizar: PropTypes.func.isRequired,
};
