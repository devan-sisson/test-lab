import { useImmer } from 'use-immer';
import './App.css'
import { FormProvider } from './FormWrapper'
import MyForm from './MyForm'

function App() {

  const validationRules = {
    notEmpty: (value) => !!value
  }

  return (
    <>
      <FormProvider
        state={useImmer()}
        validationRules={validationRules}
      >
        <MyForm />
      </FormProvider>
    </>
  )
}

export default App
