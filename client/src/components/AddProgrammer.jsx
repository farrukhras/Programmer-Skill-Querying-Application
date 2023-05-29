import { useState } from 'react';
import axios from 'axios';
import './addProgrammer.css';

function AddProgrammer() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [jobTitle, setJobTitle] = useState('');
  const [imageFile, setImageFile] = useState('');
  const [skills, setSkills] = useState([]);
  const [newSkill, setNewSkill] = useState('');

  function handleImageChange(e) {
    const file = e.target.files[0];
    setImageFile(file);
  };

  function handleSkillAdd() {
    setSkills([...skills, newSkill]);
    setNewSkill('');
  };

  async function handleSubmit(e) {
    // Convert image to a string
    const image = await readFileAsBase64(imageFile);

    // send a POST request to the backend API to create a new user
    axios.post('/users/create', { name, email, jobTitle, image, skills })
    .then(response => {
      console.log(response);
    })
    .catch(error => {
      console.log(error);
    });


    // Reset the form values
    setName('');
    setEmail('');
    setJobTitle('');
    setImageFile(null);
    setSkills([]);
    setNewSkill('');

    alert("Programmer added successfully!");
  };

  // Helper function to read the image file as a base64 string
  const readFileAsBase64 = (file) => {
    return new Promise((resolve, reject) => {
      const reader = new FileReader();
      reader.onload = () => resolve(reader.result);
      reader.onerror = (error) => reject(error);
      reader.readAsDataURL(file);
    });
  };

  return (
    <div className='formContainer'>
      <form onSubmit={handleSubmit}>
        <h1 className='addProgrammerHeading'>Add a new Programmer</h1>
        <div className='alignSameLine'>
          <label>Name:</label>
          <input type="text" value={name} onChange={(e) => {setName(e.target.value)}} required />
        </div>
        <div className='alignSameLine'>
          <label>Email:</label>
          <input type="email" value={email} onChange={(e) => {setEmail(e.target.value)}} required />
        </div>
        <div className='alignSameLine'>
          <label>Job Title:</label>
          <input type="text" value={jobTitle} onChange={(e) => {setJobTitle(e.target.value)}} required />
        </div>
        <div className='alignSameLine'>
          <label>Image:</label>
          <input type="file" accept="image/*" onChange={handleImageChange} />
        </div>
        <>
          <div className='alignSkillsItems'>
            <label>Skills:</label>
            <input type="text" className='skillInput' value={newSkill} onChange={(e) => {setNewSkill(e.target.value)}} />
            <button type="button" className='skillButton' onClick={handleSkillAdd}>Add Skill</button>
          </div>
          <div>
            <ul>
              {skills.map((skill, index) => (
                <li key={index}>{skill}</li>
              ))}
            </ul>
          </div>
        </>
        <button type="submit" className='submitButton'>Submit</button>
      </form>
    </div>
  );
}

export default AddProgrammer;