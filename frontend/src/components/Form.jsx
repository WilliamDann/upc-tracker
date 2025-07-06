import { useRef } from 'react';
import { Form, Button } from 'react-bootstrap';

const EditableObjectForm = ({ data, onSubmit }) => {
  const formRef = useRef(null);

  const handleSubmit = (e) => {
    e.preventDefault();

    const formElements = formRef.current.elements;
    const updatedData = {};

    Object.keys(data).forEach((key) => {
      const element = formElements.namedItem(key);

      if (!element) return;

      if (typeof data[key] === 'boolean') {
        updatedData[key] = element.checked;
      } else if (typeof data[key] === 'number') {
        updatedData[key] = parseFloat(element.value) || 0;
      } else {
        updatedData[key] = element.value;
      }
    });

    if (onSubmit) {
      onSubmit(updatedData);
    }
  };

  const renderInput = (key, value) => {
    const type = typeof value;

    if (type === 'boolean') {
      return (
        <Form.Check
          key={key}
          type="checkbox"
          label={key}
          defaultChecked={value}
          name={key}
          className="mb-3"
        />
      );
    }

    return (
      <Form.Group controlId={`form-${key}`} className="mb-3" key={key}>
        <Form.Label>{key}</Form.Label>
        <Form.Control
          type={type === 'number' ? 'number' : 'text'}
          defaultValue={value}
          name={key}
        />
      </Form.Group>
    );
  };

  return (
    <Form ref={formRef} onSubmit={handleSubmit}>
      {Object.entries(data).map(([key, value]) => renderInput(key, value))}
      <Button variant="primary" type="submit">
        Save
      </Button>
    </Form>
  );
};

export default EditableObjectForm;
