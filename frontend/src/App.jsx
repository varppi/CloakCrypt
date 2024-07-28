import Bar from './components/Header.jsx'
import InitMenu from './pages/InitMenu.jsx'
import Create from './pages/Create.jsx'
import Decrypt from './pages/Decrypt.jsx'

import { useState } from 'react'

function App() {
    var state = useState(-1)
    var page = <></>
    
    switch (state[0]) {
    case -1:
        page = <InitMenu state={state}></InitMenu>
        break
    case 0:
        page = <Decrypt></Decrypt>
        break
    case 1:
        page = <Create></Create>
        break
    }

    return <>
    <Bar></Bar>
    {page}
    </>
}

export default App
