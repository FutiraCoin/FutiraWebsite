
document.addEventListener('DOMContentLoaded',()=>{
    gsap.registerPlugin(ScrollTrigger)
    const iconWrapper = Array.from(document.querySelectorAll(".icon_card__icon[data-type='icon']"))
    let style = document.createElement("style")
    document.head.appendChild(style)
    const FontAwesomeMap = {
        fas :{
            size: 900,
            family:"Pro"
        },
        fab:{
            size: 400,
            family:"Brands"
        },
        far:{
            size: 400,
            family:"Pro"
        }
    }
    iconWrapper.forEach(element => {
        style.append(`
            .icon_card__icon[data-code='${element.dataset.code}']::before{
                content: '\\${element.dataset.code}';
                font-family: "Font Awesome 5 ${FontAwesomeMap[element.dataset.icon.toLowerCase()].family}";
                font-weight: ${FontAwesomeMap[element.dataset.icon.toLowerCase()].size};
            }
        `)
    });
    ScrollTrigger.defaults({
        //markers: true,
    })
    const duration = .3
    const fadeUp={y:50 , opacity:0 }
    const whatWeOfferTimeLine = gsap.timeline({scrollTrigger:{trigger : ".what_we_offer",},duration :duration,})
    const benfitsTimeLine = gsap.timeline({scrollTrigger:{trigger : ".benfits",},duration :duration,})
    const gallaryTimeLine = gsap.timeline({scrollTrigger:{trigger : ".gallary",},duration :duration,})
    const priceTimeLine = gsap.timeline({scrollTrigger:{trigger : ".price",},duration :duration,})
    const testimonialTimeLine = gsap.timeline({scrollTrigger:{trigger : ".testimonial",},duration :duration,})
    
    whatWeOfferTimeLine.from(".what_we_offer__title", fadeUp).from(".what_we_offer_card",{...fadeUp,stagger:.05,});
    benfitsTimeLine.from(".benfits__title",fadeUp).from(".benfits__text",fadeUp).from(".benfits_learn",fadeUp).from(".benfits__list",fadeUp);
    gallaryTimeLine.from(".gallary__title" , fadeUp).from(".gallary__text" ,  fadeUp).from(".gallary__video" ,  fadeUp ,"<=");
    priceTimeLine.from(".price__title",fadeUp).from(".price__text",fadeUp).from(".card-group",fadeUp);
    testimonialTimeLine.from(".testimonial__text",fadeUp).from(".carousel",fadeUp);
})