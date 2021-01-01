import * as PropTypes from 'prop-types';
import * as React from 'react';
import { DivContainer, DivRow } from './styles';
import { FormCrearPelicula } from '../../components/FormCrearPelicula';
import { ListaPeliculas } from '../../components/ListarPeliculas';
import { Pelicula } from '../../models/Pelicula';
import { useEffect } from 'react';

interface GestionPeliculasProps {
  peliculas: Array<Pelicula>;
  listarPeliculas: () => void;
  agregarNuevaPelicula: (pelicula: Pelicula) => void;
  eliminarPelicula: (pelicula: Pelicula) => void;
}

export const GestionPeliculas: React.FC<GestionPeliculasProps> = ({
  agregarNuevaPelicula,
  peliculas,
  listarPeliculas,
  eliminarPelicula,
}) => {
  useEffect(() => {
    listarPeliculas();
  }, [listarPeliculas]);
  return (
    <DivContainer>
      <DivRow>
        <FormCrearPelicula
          onSubmit={agregarNuevaPelicula}
          formTitle="Crear pelicula"
        />
      </DivRow>
      <DivRow>
        <ListaPeliculas
          peliculas={peliculas}
          onClickEliminarPelicula={eliminarPelicula}
        />
      </DivRow>
    </DivContainer>
  );
};

GestionPeliculas.propTypes = {
  peliculas: PropTypes.array.isRequired,
  listarPeliculas: PropTypes.func.isRequired,
  agregarNuevaPelicula: PropTypes.func.isRequired,
  eliminarPelicula: PropTypes.func.isRequired,
};
