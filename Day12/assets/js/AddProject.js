let dataBlog = []

function addBlog (event) {
    event.preventDefault()

    let name = document.getElementById("ProjectName").value
    let startdate = document.getElementById("StartDate").value
    let enddate = document.getElementById("EndDate").value
    let description = document.getElementById("Description").value
    let reactjs = document.getElementById("ReactJs").checked
    let javascript = document.getElementById("Javascript").checked
    let android = document.getElementById("Android").checked
    let nodejs = document.getElementById("NodeJs").checked
    let image = document.getElementById("input-file").files

    if (image[0]){
        image = URL.createObjectURL(image[0])
    }

    console.log(image)

    if (reactjs){
        reactjs = `<i class="fa-brands fa-react me-3"></i>`
    } else {
        reactjs = ``
    }
    if (javascript){
        javascript = `<i class="fa-brands fa-js me-3"></i> `
    } else {
        javascript = ``
    }
    if (android){
        android = `<i class="fa-brands fa-android me-3"></i>`
    } else {
        android = ``
    }
    if (nodejs){
        nodejs = `<i class="fa-brands fa-node-js"></i> `
    } else {
        nodejs = ``
    }

    let inputstartdate = new Date (startdate)
    let inputenddate = new Date(enddate)

    let test = ``

    let timedistance = inputenddate - inputstartdate

    let distanceseconds = Math.floor(timedistance / 1000)
    let distanceminutes = Math.floor(distanceseconds / 60)
    let distancehours = Math.floor(distanceminutes / 60)
    let distancedays = Math.floor(distancehours / 24)
    let distanceweeks = Math.floor(distancedays / 7)
    let distancemonths = Math.floor(distanceweeks / 4)
    let distanceyears = Math.floor(distancemonths / 12)

    if (distanceseconds >= 60 && distanceminutes < 60) {
        test = `${distanceminutes} menit`
    } else if (distanceminutes >= 60 && distancehours < 24) {
        test = `${distancehours} jam`
    } else if (distancehours >= 24 && distancedays < 7) {
        test = `${distancedays} hari`
    } else if (distancedays >= 7 && distanceweeks < 4) {
        test = `${distanceweeks} minggu`
    } else if (distanceweeks >= 4 && distancemonths < 12) {
        test = `${distancemonths} bulan`
    } 

    let blog = {
        name,
        startdate,
        enddate,
        description,
        test,
        reactjs,
        javascript,
        android,
        nodejs,
        image,
    }

    dataBlog.push(blog)
    renderBlog()

    console.log(dataBlog)

      if (name == "") {
        return alert("Please input your project name or title");
      } else if (startdate == "") {
        return alert("When did you start this project?");
      } else if (enddate == "") {
        return alert("When did you finish this project?");
      } else if (description == "") {
        return alert("Please describe this project.");
      } else if (!image[0]) {
        return alert("Please attach an image of your project.");
      }
}

function renderBlog() {
    document.getElementById("myproject").innerHTML = ``

    for (let index = 0; index < dataBlog.length; index++){
        document.getElementById("myproject").innerHTML += `
        <div class="row m-3">
        <div class="col-auto d-flex justify-content-center my-3">
            <div class="card card-myproject" style="width: 20rem; height: 33rem;">
              <img src="${dataBlog[index].image}" class="card-img-top object-fit-cover p-2" style="width: 100%; height: 250px;" alt="...">
              <div class="card-body">
              <a href="ProjectDetail" class="text-decoration-none text-black"
                <h5 class="card-title cardtittle">${dataBlog[index].name}</h5>
                <p class="card-duration my-1">Durasi: ${dataBlog[index].test}</p>
                <p class="cardtext my-2 overflow-hidden white-space-wrap" style="height: 7em;">${dataBlog[index].description}</p>
                <div class="icon my-2">
                    ${dataBlog[index].reactjs}                        
                    ${dataBlog[index].javascript}                      
                    ${dataBlog[index].android}     
                    ${dataBlog[index].nodejs}              
              </div>
                <a href="#" class="btn btn-dark" style="width: 49%;">Edit</a>
                <a href="#" class="btn btn-dark" style="width: 49%;">Delete</a>
            </div>
        </div>
        </div>
        `
    }
}