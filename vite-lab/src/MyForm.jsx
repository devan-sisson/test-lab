import { useForm } from "./FormWrapper";

const validationRules = {
  notEmpty: "notEmpty"
}

export default function MyForm() {
  const {formState, addField, onBlurValidate} = useForm();

  const handleFirstName = addField("firstName", "First Name must be filled in", "", validationRules.notEmpty)

  return (
    <>
      <input type="text" name="first-name" id="first-name" onChange={(e) => handleFirstName(e.target.value)} onBlur={onBlurValidate("firstName")} value={formState?.firstName} />
      <p>{formState.firstName}</p>
    </>

  )
}