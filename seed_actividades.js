const API_USUARIOS = 'http://localhost:8082';
const API_ACTIVIDADES = 'http://localhost:8081';

async function seed() {
    console.log("Haciendo login con usuario admin...");
    const loginRes = await fetch(`${API_USUARIOS}/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username: "admin", password: "admin" })
    });

    if (!loginRes.ok) throw new Error("Login failed");
    const { token, id: owner_id } = await loginRes.json();
    console.log("Token obtenido. Owner ID:", owner_id);

    const actividades = [
        {
            nombre: "Crossfit",
            descripcion: "Entrenamiento de alta intensidad con ejercicios funcionales constantemente variados. Requiere experiencia previa.",
            profesor: "Carlos Fit",
            owner_id,
            horarios: [
                { dia: "Lunes", horarioInicio: "08:00", horarioFinal: "09:00", cupo: 20 },
                { dia: "Miércoles", horarioInicio: "08:00", horarioFinal: "09:00", cupo: 20 },
                { dia: "Viernes", horarioInicio: "08:00", horarioFinal: "09:00", cupo: 20 }
            ]
        },
        {
            nombre: "Yoga Vinyasa",
            descripcion: "Flujo dinámico de posturas sincronizadas con la respiración. Ideal para relajarse y mejorar firmeza.",
            profesor: "Laura Zen",
            owner_id,
            horarios: [
                { dia: "Martes", horarioInicio: "19:00", horarioFinal: "20:30", cupo: 15 },
                { dia: "Jueves", horarioInicio: "19:00", horarioFinal: "20:30", cupo: 15 }
            ]
        },
        {
            nombre: "Spinning",
            descripcion: "Clase de ciclismo indoor con música motivadora y alta quema de calorías.",
            profesor: "Martín Rueda",
            owner_id,
            horarios: [
                { dia: "Lunes", horarioInicio: "18:00", horarioFinal: "19:00", cupo: 25 },
                { dia: "Miércoles", horarioInicio: "18:00", horarioFinal: "19:00", cupo: 25 }
            ]
        },
        {
            nombre: "Pilates Reformer",
            descripcion: "Mejora tu postura, flexibilidad y fuerza core usando equipos y camillas especializadas.",
            profesor: "Ana Flex",
            owner_id,
            horarios: [
                { dia: "Lunes", horarioInicio: "10:00", horarioFinal: "11:00", cupo: 10 },
                { dia: "Miércoles", horarioInicio: "10:00", horarioFinal: "11:00", cupo: 10 }
            ]
        },
        {
            nombre: "Zumba Fitness",
            descripcion: "Baila al ritmo de la música latina y quema calorías divirtiéndote sin darte cuenta.",
            profesor: "Sonia Ritmo",
            owner_id,
            horarios: [
                { dia: "Viernes", horarioInicio: "20:00", horarioFinal: "21:00", cupo: 30 }
            ]
        },
        {
            nombre: "Natación Libre",
            descripcion: "Acceso libre a la piscina climatizada semi-olímpica con carriles divididos por nivel.",
            profesor: "Juan Aguas",
            owner_id,
            horarios: [
                { dia: "Lunes", horarioInicio: "09:00", horarioFinal: "11:00", cupo: 15 },
                { dia: "Sábado", horarioInicio: "09:00", horarioFinal: "13:00", cupo: 15 }
            ]
        },
        {
            nombre: "Boxeo Recreativo",
            descripcion: "Aprende técnica de boxeo, mejora tus reflejos y descarga estrés en los sacos.",
            profesor: "Tyson Junior",
            owner_id,
            horarios: [
                { dia: "Martes", horarioInicio: "21:00", horarioFinal: "22:00", cupo: 12 },
                { dia: "Jueves", horarioInicio: "21:00", horarioFinal: "22:00", cupo: 12 }
            ]
        },
        {
            nombre: "Entrenamiento Funcional",
            descripcion: "Circuitos dinámicos para trabajar todo el cuerpo mejorando fuerza, velocidad y resistencia.",
            profesor: "Carlos Fit",
            owner_id,
            horarios: [
                { dia: "Lunes", horarioInicio: "07:00", horarioFinal: "08:00", cupo: 20 },
                { dia: "Miércoles", horarioInicio: "07:00", horarioFinal: "08:00", cupo: 20 },
                { dia: "Viernes", horarioInicio: "07:00", horarioFinal: "08:00", cupo: 20 }
            ]
        },
        {
            nombre: "Stretching",
            descripcion: "Relaja tus músculos y previene lesiones con estiramientos profundos y control postural.",
            profesor: "Laura Zen",
            owner_id,
            horarios: [
                { dia: "Lunes", horarioInicio: "17:00", horarioFinal: "18:00", cupo: 15 }
            ]
        },
        {
            nombre: "Halterofilia",
            descripcion: "Centro especializado con plataformas y barras olímpicas para levantamiento de pesas.",
            profesor: "Arnold Max",
            owner_id,
            horarios: [
                { dia: "Martes", horarioInicio: "14:00", horarioFinal: "16:00", cupo: 8 },
                { dia: "Jueves", horarioInicio: "14:00", horarioFinal: "16:00", cupo: 8 }
            ]
        }
    ];

    for (const act of actividades) {
        console.log(`Intentando crear actividad: ${act.nombre}...`);
        try {
            const res = await fetch(`${API_ACTIVIDADES}/actividad`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify(act)
            });
            if (!res.ok) {
                const err = await res.text();
                console.error(`❌ Error creando ${act.nombre}:`, err);
            } else {
                console.log(`✅ OK: ${act.nombre} creada.`);
            }
        } catch (e) {
            console.error(`❌ Excepción en red con ${act.nombre}:`, e.message);
        }
    }

    console.log("Seeding completado con éxito! 🎉");
}

seed().catch(console.error);
