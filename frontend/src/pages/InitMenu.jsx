import '../css/InitMenu.css';

export default function({state}) {
    return <div>
        <h1 className='text-center title'>Welcome to CloakCrypt!</h1>
        <div className='d-inline-flex menu'>
            <button onClick={()=>{state[1](0)}}><h3><i className='bi bi-file-earmark-lock'></i> Open Encrypted Container</h3></button>
            <button onClick={()=>{state[1](1)}}><h3><i className='bi bi-file-plus'></i> Create Encrypted Container</h3></button>
        </div>
    </div>
}