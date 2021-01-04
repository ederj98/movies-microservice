import * as PropTypes from 'prop-types';
import * as React from 'react';
import { DivContainer } from '../../../../shared/components/Div/index';
import { ListaPeliculas } from '../../components/ListarPeliculas/index';
import { Pelicula } from '../../../Pelicula/models/Pelicula';
import { useEffect } from 'react';

interface GestionListarPeliculasProps {
  peliculas: Array<Pelicula>;
  listarPeliculas: () => void;
  eliminarPelicula: (pelicula: Pelicula) => void;
}

export const GestionListarPeliculas: React.FC<GestionListarPeliculasProps> = ({
  peliculas,
  listarPeliculas,
  eliminarPelicula,
}) => {
  useEffect(() => {
    listarPeliculas();
  }, [listarPeliculas]);
  return (
    <DivContainer>
        <ListaPeliculas
          peliculas={peliculas}
          onClickEliminarPelicula={eliminarPelicula}
        />
    </DivContainer>
  );
};

GestionListarPeliculas.propTypes = {
  peliculas: PropTypes.array.isRequired,
  listarPeliculas: PropTypes.func.isRequired,
  eliminarPelicula: PropTypes.func.isRequired,
};
