import { useState } from 'react'
import { SelectFile, EncryptFile } from '../../wailsjs/go/main/App'
import '../css/Form.css'

export default function() {
    var [path, setPath] = useState("")
    var [path2, setPath2] = useState("")
    var [password, setPassword] = useState("")
    var [submitted, setSubmit] = useState(null)
    var [password2, setPassword2] = useState("")
    var [buttonDisabled, setButtonDisabled] = useState(false)
    var [mainButton, setMainButton] = useState("Encrypt file")

    function submit(path, password) {
        setMainButton("Encrypting...")
        setButtonDisabled(true)
        EncryptFile(path, password).then((result) => {
            setSubmit(result)
        })
    }

    var createFile = <div className='basicform'>
    <div>
        <h3 style={{"fontWeight": "700"}}>Create Encrypted Container:</h3>
        {(password2.length > 0 && password != password2) ? <><i className="bi bi-exclamation-circle"></i> <b>Passwords don't match!</b><br></br></>: null}
        {(password2.length > 0 && password.length <= 7) ? <><i className="bi bi-exclamation-circle"></i> <b>Password should be over 7 characters long!</b><br></br></>: null}
        {(password.length == 0) ? <> <i className="bi bi-exclamation-circle"></i> <b>Enter a password</b><br></br></> : null}
        <label>Enter password:</label>
        <input type='password' className='form-control' onChange={(event)=>{setPassword(event.target.value)}}></input>        
        <label>Repeat encryption password:</label>
        <input type='password' className='form-control' onChange={(event)=>{setPassword2(event.target.value)}}></input>        
        <label>File to hide the target file in:</label>
        <button className='form-control w-50' onClick={()=>{SelectFile(0).then((path_)=>{setPath(path_)})}}>{path.length > 0 ? "Current File: "+path : "Select File"}</button>
        <label>File to encrypt:</label>
        <button className='form-control w-50' onClick={()=>{SelectFile(1).then((path_)=>{setPath2(path_)})}}>{path2.length > 0 ? "Current File: "+path2 : "Select File"}</button>
        <br></br>
        <button id='createButton' className='form-control' onClick={()=>{submit(password)}} {...(buttonDisabled ? {disabled: "true"} : {})} {...(!(password==password2&&password.length>7&&path.length>0&&path2.length>0) ? {hidden: "true"} : {})}>{mainButton}</button>
    </div>
</div>

    var submittedS = <div className='d-flex w-100 justify-content-center'>
    {
    !submitted ?
    <h1 style={{"fontSize": "50px", "color": "lightcoral"}}>Something went wrong!</h1>:
    <h1 style={{"fontSize": "50px"}}>File encrypted successfully!</h1>
    }
    </div>

    return <>
       {submitted!=null ? submittedS: createFile}
    </>
}