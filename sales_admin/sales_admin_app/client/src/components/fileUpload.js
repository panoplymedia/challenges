import React, { useState, useRef, useEffect } from "react";
import { withRouter } from "react-router-dom";

function FileUpload(props) {
  const useForceUpdate = () => useState()[1];
  const fileInput = useRef(null);
  const forceUpdate = useForceUpdate();

  useEffect(e => {
    window.addEventListener("keyup", clickFileInput);
    return () => window.removeEventListener("keyup", clickFileInput)
  })

  const clickFileInput = (e) => {
    if (fileInput.current.nextSibling.contains(document.activeElement)) {
      if (e.keyCode === 32) {
        fileInput.current.click();
      }
    }
  }

  const uploadFile = (e) => {
    e.preventDefault();
    var form = document.forms.namedItem("fileForm");
    var formData = new FormData();
    formData.append("file", form[0].files[0]);
    postFile(formData);
    window.location.reload()
  };

  const fileNames = (e) => {
    const { current } = fileInput;
      if (current && current.files.length > 0) {
        let messages = [];
        for (let file of current.files) {
          messages = messages.concat(<p key={file.name}>{file.name}</p>);
        }
         return messages;
       }
      return null;
  }

  return (
    <div className="input-group">
      <form onSubmit={uploadFile} id="fileForm" name="fileName">
          <input
            type="file"
            ref={fileInput}
            className="custom-file-input"
            id="file"
            aria-describedby="inputGroupFileAddon01"
            onChange={forceUpdate}
            multiple
          />
          <label className="custom-file-label" htmlFor="file">
            Choose file
          </label>
          {fileNames()}
          <button
            className="input-group-text"
            id="inputGroupFileAddon01"
            type="submit"
          >
            Upload
          </button>
      </form>
    </div>
  );
}

async function postFile(formData) {
  try {
    const response = await fetch("http://localhost:3001/api/upload_csv", {
      method: "POST",
      mode: "cors",
      headers: {
        Authorization: JSON.parse(window.localStorage.userAuth).token,
        type: "bearer",
      },
      body: formData
    });
    return await response.json()
  } catch (e) {
      console.log(e)
  }
} 

export default withRouter(FileUpload)