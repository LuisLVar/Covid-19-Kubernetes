import React, { useState, useEffect } from "react";
import { CanvasJSChart } from 'canvasjs-react-charts';

export const GRAFSTATE = ( { socket }) => {
    const [state, setState] = useState([]);

    useEffect(() => {

        socket.on("grafState", data =>  {

            let total = 0;
            data.forEach( element => { total = total + element.count; });

            let aux = []
            data.forEach( element => {  
                const ope = (element.count/total)*100;
                aux.push({ y: parseInt( ope.toFixed(2) , 10) , label: element._id }) 
            });
            setState(aux);
        } );
        
    }, []);

    const optionsStation = {
        exportEnabled: true,
        animationEnabled: true,
        title: {
            text: "Porcentaje por Estado"
        },
        data: [{
            type: "pie",
            startAngle: 75,
            toolTipContent: "<b>{label}</b>: {y}%",
            showInLegend: "true",
            legendText: "{label}",
            indexLabelFontSize: 16,
            indexLabel: "{label} - {y}%",
            dataPoints: state
        }]
    }

    return (
        <div className="col-md-10 offset-1 mt-4 mb-4" >
            <CanvasJSChart options={ optionsStation } />
        </div>
    );
}
