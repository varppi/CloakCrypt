import { useState } from 'react'
import { SelectFile, DecryptFile } from '../../wailsjs/go/main/App'

import '../css/Form.css'

export default function() {
    var [path, setPath] = useState("")
    var [password, setPassword] = useState("")
    var [submitted, setSubmit] = useState(null)
    var [buttonDisabled, setButtonDisabled] = useState(false)
    var [mainButton, setMainButton] = useState("Decrypt file")

    function submit(password) {
        setMainButton("Decrypting...")
        setButtonDisabled(true)
        DecryptFile(password).then((result) => {
            console.log(result)
            setSubmit(result)
        })
    }

    var decryptScreen = <div className='basicform'>
        <div>
            <h3 style={{"fontWeight": "700"}}>Open Encrypted Container:</h3>
            {(password.length == 0) ? <> <i className="bi bi-exclamation-circle"></i> <b>Enter a password</b><br></br></> : null}
            <label>Enter password:</label>
            <input type='password' className='form-control' onChange={(event)=>{setPassword(event.target.value)}}></input>        
            <label>Container to decrypt:</label>
            <button className='form-control w-50' onClick={()=>{SelectFile(2).then((path_)=>{setPath(path_)})}}>{path.length > 0 ? "Current File: "+path : "Select File"}</button>
            <br></br>
            <button id='createButton' className='form-control' onClick={()=>{submit(password)}} {...(buttonDisabled ? {disabled: "true"} : {})} {...(!(password.length>0&&path.length>0) ? {hidden: "true"} : {})}>{mainButton}</button>
        </div>
    </div>

    var submittedS = <div className='d-flex w-100 justify-content-center'>
    {
    !submitted ?
    <h1 style={{"fontSize": "50px", "color": "lightcoral"}}>Something went wrong!</h1>:
    <h1 style={{"fontSize": "50px"}}>File decrypted successfully!</h1>
    }
    </div>

    return  <>
       {submitted!=null ? submittedS: decryptScreen}
    </>
}