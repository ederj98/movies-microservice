import * as PropTypes from 'prop-types';
import * as React from 'react';
import * as Yup from 'yup';
import { Button } from '../../../../shared/components/Button';
import { FormikHelpers } from 'formik/dist/types';
import { Input } from '../../../../shared/components/Input';
import { Pelicula } from '../../models/Pelicula';
import { SpanError } from './styles';
import { useFormik } from 'formik';

interface FormValues {
  name: string;
  director: string;
  writer: string;
  stars: string
}

interface FormCrearPeliculaProp {
  onSubmit: (payload: Pelicula) => any;
  disabled?: boolean;
  formTitle: string;
  initialValues?: FormValues;
}

const validationSchema = Yup.object().shape<FormValues>({
  name: Yup.string().required('El campo Name es requerido.'),
  director: Yup.string().required('El campo Director es requerido.'),
  writer: Yup.string().required('El campo Writer es requerido.'),
  stars: Yup.string().required('El campo Stars es requerido.'),
});

export const FormCrearProducto: React.FC<FormCrearPeliculaProp> = ({
  onSubmit,
  disabled,
  formTitle,
  initialValues = {
    name: '',
    director: '',
    writer: '',
    stars: '',
  },
}) => {
  const handleSubmit = (
    values: FormValues,
    { resetForm }: FormikHelpers<FormValues>
  ) => {
    onSubmit({
      Id: '',
      Name: values.name,
      Director: values.director,
      Writer: values.writer,
      Stars: values.stars,
    });
    resetForm();
  };
  const formik = useFormik({
    initialValues,
    validationSchema,
    onSubmit: handleSubmit,
  });

  return (
    <form onSubmit={formik.handleSubmit}>
      <h2>{formTitle}</h2>
      <Input
        disabled={disabled}
        name="name"
        placeholder="Name"
        value={formik.values.name}
        onChange={formik.handleChange}
      />
      {formik.touched.name && formik.errors.name && (
        <SpanError>{formik.errors.name}</SpanError>
      )}
      <Input
        disabled={disabled}
        name="director"
        placeholder="Director"
        value={formik.values.director}
        onChange={formik.handleChange}
      />
      {formik.touched.director && formik.errors.director && (
        <SpanError>{formik.errors.director}</SpanError>
      )}
      <Input
        disabled={disabled}
        name="writer"
        placeholder="Writer"
        value={formik.values.writer}
        onChange={formik.handleChange}
      />
      {formik.touched.writer && formik.errors.writer && (
        <SpanError>{formik.errors.writer}</SpanError>
      )}
      <Input
        disabled={disabled}
        name="stars"
        placeholder="Stars"
        value={formik.values.stars}
        onChange={formik.handleChange}
      />
      {formik.touched.stars && formik.errors.stars && (
        <SpanError>{formik.errors.stars}</SpanError>
      )}
      <Button type="submit">Registrar</Button>
    </form>
  );
};

FormCrearProducto.propTypes = {
  onSubmit: PropTypes.func.isRequired,
  formTitle: PropTypes.string.isRequired,
  disabled: PropTypes.bool,
  initialValues: PropTypes.shape({
    name: PropTypes.string.isRequired,
    director: PropTypes.string.isRequired,
    writer: PropTypes.string.isRequired,
    stars: PropTypes.string.isRequired,
  }),
};
