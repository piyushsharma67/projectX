import { TextField } from "@mui/material"

interface IInputCompProps {

}
function InputComp(props: IInputCompProps) {
    return (
        <TextField
            {...props}
        />
    )
}

export default InputComp