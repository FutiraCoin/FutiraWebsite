
document.addEventListener("DOMContentLoaded",function() {
    gsap.registerPlugin(Flip);
    const projects = gsap.utils.toArray(".project");
    const filters = document.querySelectorAll(".projects__filters li button");
    filters.forEach(filter=>{
        filter.addEventListener("click",_ =>{
            document.querySelector(".selected").classList.remove("selected");
            filter.classList.add("selected");
            updateFilters(filter.dataset.filter)
        })
    })
    function updateFilters(filter) {
        const state = Flip.getState([projects]);
        const duration = .4;
        /*let matched = 0;
        let eleHeight = 0;
        const columnCount = 4;
        */
        projects.forEach(project=>{
            /*if(project.dataset.filter === filter ||filter === "0" ){
                project.style.display ="inline-flex";
                //matched++
            }else{
                project.style.display ="none";
            }
            if( eleHeight === 0 ){
                eleHeight = project.offsetHeight
            }*/
            project.style.display =(project.dataset.filter === filter ||filter === "0" )?"inline-flex" :"none" ;
        })
        
        Flip.from(state,{
            duration:duration,
            absolute:true,
            scale:true,
            ease: "power1.inOut",
            onEnter:Element => gsap.fromTo(Element,{scale:0 , opacity:0},{scale:1 , opacity:1 , duration:duration ,}),
            onLeave : Element=> gsap.to(Element, {scale:0 , opacity:0 ,duration:duration}),
            //onStart : _=>gsap.to("#project-holder",{height:Math.ceil(matched/ columnCount) * eleHeight , duration:.3 ,ease: "power1.inOut",})
        })
    }   
})