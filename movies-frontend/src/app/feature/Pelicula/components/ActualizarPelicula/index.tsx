import { func, string as str, bool } from 'prop-types';
import { object, number, string } from 'yup';
import { Button } from '../../../../shared/components/Button';
import { FormikHelpers } from 'formik/dist/types';
import { Input } from '../../../../shared/components/Input';
import { Pelicula } from '../../models/Pelicula';
import { SpanError } from './styles';
import { useFormik } from 'formik';
import React, { useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';

interface FormValues {
  id: number
  name: string;
  director: string;
  writer: string;
  stars: string
}

interface FormActualizarPeliculaProp {
  pelicula: Pelicula;
  onSubmit: (payload: Pelicula) => any;
  disabled?: boolean;
  formTitle: string;
}

const validationSchema = object().shape<FormValues>({
  id: number().required('El campo Id es requerido.'),
  name: string().required('El campo Name es requerido.'),
  director: string().required('El campo Director es requerido.'),
  writer: string().required('El campo Writer es requerido.'),
  stars: string().required('El campo Stars es requerido.'),
});

export const FormActualizarPelicula: React.FC<FormActualizarPeliculaProp> = ({
  pelicula,
  onSubmit,
  disabled,
  formTitle,
}) => {
  const [initialValues, setInitialValues] = useState({
    id: pelicula.id || 0,
    name: pelicula.name || '',
    director: pelicula.director || '',
    writer: pelicula.writer || '',
    stars: pelicula.stars || '',
  })
  useEffect(() => {
    setInitialValues(pelicula)
  }, [pelicula, initialValues]);
  
  let history = useHistory()
  const handleSubmit = (
    values: FormValues,
    { resetForm }: FormikHelpers<FormValues>
  ) => {
    onSubmit({
      id: pelicula.id,
      name: values.name,
      director: values.director,
      writer: values.writer,
      stars: values.stars,
    });
    resetForm();
    history.goBack()
    //window.location.href = "/home";
  };
  const formik = useFormik({
    initialValues,
    validationSchema,
    onSubmit: handleSubmit,
    enableReinitialize: true,
  });

  return (
    <form onSubmit={formik.handleSubmit}>
      <h2>{formTitle}</h2>
      <br/>
      {formik.touched.name && formik.errors.name && (
        <SpanError>{formik.errors.name}</SpanError>
      )}
      <Input
        disabled={disabled}
        name="name"
        placeholder="Name"
        value={formik.values.name}
        onChange={formik.handleChange}
      />
      {formik.touched.director && formik.errors.director && (
        <SpanError>{formik.errors.director}</SpanError>
      )}
      <Input
        disabled={disabled}
        name="director"
        placeholder="Director"
        value={formik.values.director}
        onChange={formik.handleChange}
      />
      {formik.touched.writer && formik.errors.writer && (
        <SpanError>{formik.errors.writer}</SpanError>
      )}
      <Input
        disabled={disabled}
        name="writer"
        placeholder="Writer"
        value={formik.values.writer}
        onChange={formik.handleChange}
      />
      {formik.touched.stars && formik.errors.stars && (
        <SpanError>{formik.errors.stars}</SpanError>
      )}
      <Input
        disabled={disabled}
        name="stars"
        placeholder="Stars"
        value={formik.values.stars}
        onChange={formik.handleChange}
      />
      <br/>
      <Button type="submit">Actualizar</Button>
    </form>
  );
};

FormActualizarPelicula.propTypes = {
  onSubmit: func.isRequired,
  formTitle: str.isRequired,
  disabled: bool,
};
