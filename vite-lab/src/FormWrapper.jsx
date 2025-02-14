/* eslint-disable react/prop-types */
import { createContext, useCallback, useContext } from "react";
import { useImmer } from "use-immer";

const FormContext = createContext();

function validateField(draftErrors, record, field, isValidFn, errorMessage, theFieldToValidate) {
  const isValid = isValidFn(record[field], record);
  const shouldValidate = !theFieldToValidate || (field === theFieldToValidate);

  if (shouldValidate) {
    if (isValid) {
      delete draftErrors[field];
    } else {
      draftErrors[field] = {
        message: errorMessage,
        fieldId: field,
      };
    }
  }
}

export function FormProvider(props) {
  const { state, children, validationRules } = props;
  const [formState, setFormState] = state;
  const [formErrors, setFormErrors] = useImmer({});
  const [validations, setValidations] = useImmer([]);

  const updateField = useCallback((field) => (value) => setFormState(draft => draft[field] = value), [setFormState])

  const addField = (field, message, value, type) => {
    setValidations(draft => {
      draft.push(...draft, { field, message, validate: validationRules[type] })
    })
    setFormState(draft => {
      draft[field] = value
    })

    return updateField(field)
  }

  const validateForm = useCallback(function _validateForm(state, previousFormErrors, fieldToValidate) {
    const errors = fieldToValidate ? { ...previousFormErrors } : /** @type {FormErrors} */ ({});

    validations.forEach(v => validateField(errors, state, v.field, v.validate, v.message, fieldToValidate))

    return { errors, isValid: !Object.values(errors).length };
  }, [validations])

  const formErrorsItems = Object.entries(formErrors);
  const onBlurValidate = useCallback(
    (field) => () => setFormErrors(validateForm(formState, formErrors, field).errors),
    [formState, formErrors, setFormErrors, validateForm]
  );

  const value = {
    formState,
    addField,
    onBlurValidate
  }

  return (<>
    <FormContext.Provider value={value}>
      {children}
    </FormContext.Provider>
    {
      formErrorsItems.length
        ? (
          <div className="info-box mb-spacing">

            <span><strong>Some errors have been found:</strong></span>
            <ul>
              {
                formErrorsItems.map(([field, error]) => (
                  <li key={`form-example__error__${field}`}>
                    <a
                      href={`#${error.fieldId}`}
                      onClick={(e) => {
                        e.preventDefault();
                        const input = document.getElementById(error.fieldId);
                        input?.focus();

                        const label = input?.closest('.input-wrapper')?.querySelector('label');
                        label?.scrollIntoView();
                      }}
                    >
                      {error.message}
                    </a>
                  </li>
                ))
              }
            </ul>
          </div>
        )
        : null
    }
  </>
  )
}

export function useForm() {
  return useContext(FormContext);
}
