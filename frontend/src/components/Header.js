import logo from '../icons/graph-user-logo.jpg'
import '../css/Header.css'

export default function() {
  return (
    <div className="navbar">
      <img alt="logo" className="app-logo" src={logo} />
      <span className="greet-user">Welcome, <u>Manoj Saravanan</u></span>
    </div>
  )
}