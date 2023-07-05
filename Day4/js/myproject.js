let dataBlog = []

const addBlog = (event) => {
    event.preventDefault();

    let project = document.getElementById("input-name").value
    let startdate = document.getElementById("start-date").value
    let enddate = document.getElementById("end-date").value
    let description = document.getElementById("description").value
    let reactjs = document.getElementById("input-reactjs").checked
    let javascript = document.getElementById("input-javascript").checked
    let android = document.getElementById("input-android").checked
    let nodejs = document.getElementById("input-nodejs").checked
    let image = document.getElementById("input-file").files

    // image = URL.createObjectURL(image[0])

        
    let blog = {
        project,
        startdate,
        enddate,
        description,
        image,
    }
    
    dataBlog.push(blog)
    // renderBlog()

    console.log(dataBlog)

    if (project == "") {
        return alert("Please input your project name or title");
      } else if (startdate == "") {
        return alert("When did you start this project?");
      } else if (enddate == "") {
        return alert("When did you finish this project?");
      } else if (description == "") {
        return alert("Please describe this project.");
      } else if (image == "") {
        return alert("Please attach an image of your project.");
      }

}

// function renderBlog() {
//     document.getElementById("contents").innerHTML = ''

//     for (let index = 0; index > dataBlog.length; index++) {
//         document.getElementById("contents").innerHTML += `<div class="project">
//         <div class="project1">
//                 <img src="${dataBlog[index].image}" alt=""/>
//                 <h3>Aku Hobi Main Bola - 2021</h3>
//                 <h5>Durasi: 3 bulan</h5>
//                 <h5>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Laborum cum aut doloribus dolorum explicabo voluptas quos nam sequi, quam assumenda ratione molestiae doloremque quibusdam laboriosam iusto rem natus facilis totam.</h5>
//                 <div class="icon">
//                     <i class="fa-brands fa-react"></i>                        
//                     <i class="fa-brands fa-js"></i>                       
//                     <i class="fa-brands fa-android"></i>     
//                     <i class="fa-brands fa-node-js"></i>               
//                 </div>    
//                 <div class="btn">
//                     <button>Edit</button>
//                     <button>Delete</button>
//                 </div>
//         </div>
//     </div>
//     `
//     }
// }







// punya org

    // function emptyformalert () {
    //     let project = document.getElementById("input-name").value
    //     let startdate = document.getElementById("start-date").value
    //     let enddate = document.getElementById("end-date").value
    //     let description = document.getElementById("description").value
    //     let multicok = document.querySelectorAll("multicok:checked")
    //     let image = document.getElementById("input-file").files



    // if (project == "") {
    //     return alert("isi dong namanya!");
    //   } else if (startdate == "") {
    //     return alert("mulainya kapan?");
    //   } else if (enddate == "") {
    //     return alert("selesainya kapan?");
    //   } else if (description == "") {
    //     return alert("dijelasin dong");
    //   } else if (multicok.length === 0) {
    //     return alert("pilih dulu");
    //   } else if (image == "") {
    //     return alert("isi dulu");
    //   }
    // }

    // let dataBlog = []

    // function addBlog(event) {
    //     event.preventDefault()

    //     let project = document.getElementById("input-name").value
    //     let startdate = document.getElementById("start-date").value
    //     let enddate = document.getElementById("end-date").value
    //     let description = document.getElementById("description").value
    //     let image = document.getElementById("input-file").files

    //     const reactjsicon = '<i class="fa-brands fa-react"></i>'
    //     const javascripticon = '<i class="fa-brands fa-js"></i>'
    //     const androidicon = '<i class="fa-brands fa-android"></i>'
    //     const nodejsicon = '<i class="fa-brands fa-node-js"></i>'

    //     let reactjs = document.getElementById("input-reactjs").checked ? reactjsicon : ""
    //     let javascript = document.getElementById("input-javascript").checked ? javascripticon : ""
    //     let android = document.getElementById("input-android").checked ? androidicon : ""
    //     let nodejs = document.getElementById("input-nodejs").checked ? nodejsicon : ""

    //     let multicok = document.querySelectorAll(".multicok:checked")
    //     if (multicok.length === 0) {
    //         return alert("pilih satu dong")
    //     }

    //     image = URL.createObjectURL(image[0])
    //     console.log(image)

    //     let inputstartdate = new Date(startdate)
    //     let inputenddate = new Date(enddate)

    //     if (inputstartdate > inputenddate) {
    //         return alert("tanggalnya donggg")
    //     }

    //     let data = {
    //         project,
    //         startdate,
    //         enddate,
    //         description,
    //         reactjs,
    //         javascript,
    //         android,
    //         nodejs,
    //         image,
    //     }

    //     dataBlog.push(data)
    //     console.log(dataBlog)


    // }