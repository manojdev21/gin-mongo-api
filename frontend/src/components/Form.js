import '../css/Form.css'

export default function() {
  return (
    <div className="form-container">
      <form id="form">
        <label className="title">Full Name</label>
        <input name="fullName" className="field" />
        <label className="title">Email</label>
        <input name="email" className="field" type="email" />
        <label className="title address">Address</label>
        <div id="sub">
          <label className="title">Door Number</label>
          <input name="doorNo" className="field" />
          <label className="title">Street</label>
          <input name="street" className="field" />
          <label className="title">City</label>
          <input name="city" className="field" />
          <label className="title">Postal Code</label>
          <input name="postalCode" className="field" type="number" />
        </div>
        <button className="update-btn btn">Update</button>
        <button className="reset-btn btn" type="reset">Cancel</button>
      </form>
    </div>
  )
}