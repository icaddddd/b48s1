let dataBlog = []

let addBlog = (event) => {
    event.preventDefault()

    let project = document.getElementById("input-name").value
    let startdate = document.getElementById("start-date").value
    let enddate = document.getElementById("end-date").value
    let description = document.getElementById("description").value
    let nodejs = document.getElementById("input-nodejs").checked
    let nextjs = document.getElementById("input-nextjs").checked
    let reactjs = document.getElementById("input-reactjs").checked
    let typescript = document.getElementById("input-typescript").checked
    let image = document.getElementById("input-file").files

    image = URL.createObjectURL(image[0])

    let blog = {
        project,
        startdate,
        enddate,
        description,
        image,
    }
    
    dataBlog.push(blog)
    renderBlog()

    console.log(dataBlog)
    console.log(project)
}




    

