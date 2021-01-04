import { EstadoPelicula } from '../../modelo/EstadoPelicula';
import { Pelicula } from '../../../../feature/Pelicula/models/Pelicula';
import { agregarNuevaPelicula, eliminarPelicula,buscarPelicula, actualizarPelicula } from '../../acciones/peliculas/PeliculasAcciones';
import reductorPeliculas from './peliculasReductor';

describe('Reductor peliculas', () => {
  it('debería agregar peliculas', () => {
    // Arrange
    const estadoInicial: EstadoPelicula = {
      peliculas: [],
      pelicula: {Id: 0, Name: '', Director: '', Writer: '', Stars: ''}
    };
    const nuevaPelicula: Pelicula = {
        Id: 1,
        Name: 'Stargate',
        Director: 'Eder',
        Writer: 'The grinch',
        Stars: 'John, Prim, Katniss',
    };
    const estadoEsperado: EstadoPelicula = {
      ...estadoInicial,
      peliculas: [nuevaPelicula],
    };

    // Act
    const nuevoEstado = reductorPeliculas(
      estadoInicial,
      agregarNuevaPelicula(nuevaPelicula)
    );

    // Assert
    expect(nuevoEstado).toStrictEqual(estadoEsperado);
  });

  it('debería eliminar pelicula', () => {
    // Arrange
    const nuevaPelicula: Pelicula = {
        Id: 1,
        Name: 'Stargate',
        Director: 'Eder',
        Writer: 'The grinch',
        Stars: 'John, Prim, Katniss',
    };
    const estadoInicial: EstadoPelicula = {
      peliculas: [nuevaPelicula],
      pelicula: {Id: 0, Name: '', Director: '', Writer: '', Stars: ''}
    };
    const estadoEsperado: EstadoPelicula = {
      ...estadoInicial,
      peliculas: [],
    };

    // Act
    const nuevoEstado = reductorPeliculas(
      estadoInicial,
      eliminarPelicula(nuevaPelicula)
    );

    // Assert
    expect(nuevoEstado).toStrictEqual(estadoEsperado);
  });

  it('debería buscar pelicula', () => {
    // Arrange
    const estadoInicial: EstadoPelicula = {
      peliculas: [],
      pelicula: {Id: 0, Name: '', Director: '', Writer: '', Stars: ''}
    };
    const nuevaPelicula: Pelicula = {
      Id: 1,
      Name: 'Stargate',
      Director: 'Eder',
      Writer: 'The grinch',
      Stars: 'John, Prim, Katniss',
  };
    const estadoEsperado: EstadoPelicula = {
      ...estadoInicial,
      pelicula: nuevaPelicula,
    };

    // Act
    const nuevoEstado = reductorPeliculas(
      estadoInicial,
      buscarPelicula(nuevaPelicula)
    );

    // Assert
    expect(nuevoEstado).toStrictEqual(estadoEsperado);
  });

  it('debería actualizar pelicula', () => {
    // Arrange
    const nuevaPelicula: Pelicula = {
      Id: 1,
      Name: 'Stargate',
      Director: 'Eder',
      Writer: 'The grinch',
      Stars: 'John, Prim, Katniss',
    };
    const estadoInicial: EstadoPelicula = {
      peliculas: [nuevaPelicula],
      pelicula: {Id: 0, Name: '', Director: '', Writer: '', Stars: ''}
    };
    const estadoEsperado: EstadoPelicula = {
      ...estadoInicial,
      peliculas: [nuevaPelicula],
    };

    // Act
    const nuevoEstado = reductorPeliculas(
      estadoInicial,
      actualizarPelicula(nuevaPelicula)
    );

    // Assert
    expect(nuevoEstado).toStrictEqual(estadoEsperado);
  });
});
