import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from './ui/card';
import { Button } from './ui/button';
import { Badge } from './ui/badge';
import { Input } from './ui/input';
import { Clock, User, Calendar, Trash2, Edit, Save, X, Users } from 'lucide-react';

function ListadoActividades({ filtro, refrescar, esAdmin }) {
  const [actividades, setActividades] = useState([]);
  const [editando, setEditando] = useState(null);
  const [formData, setFormData] = useState({ nombre: '', descripcion: '', profesor: '', horarios: [] });
  const navigate = useNavigate();

  const fetchActividades = async () => {
    try {
      const res = await fetch(`${process.env.REACT_APP_API_BUSQUEDAS_URL}/search/actividades?nombre=${filtro || ''}`);
      if (!res.ok) throw new Error("Error en la respuesta del servidor");
      const data = await res.json();
      setActividades(data.actividades || []);
    } catch (error) {
      console.error("Error al obtener actividades:", error);
      setActividades([]);
    }
  };

  useEffect(() => {
    fetchActividades();
  }, [filtro, refrescar]);

  const handleEliminar = async (id) => {
    if (!window.confirm("¿Estás seguro de que deseas eliminar esta actividad?")) return;

    try {
      const tokenCookie = document.cookie.split('; ').find(row => row.startsWith('token='));
      const token = tokenCookie ? tokenCookie.split('=')[1] : null;

      const res = await fetch(`${process.env.REACT_APP_API_ACTIVIDADES_URL}/actividad/${id}`, {
        method: 'DELETE',
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });

      if (res.ok) {
        alert("Actividad eliminada correctamente");
        fetchActividades();
      } else {
        alert("Error al eliminar la actividad");
      }
    } catch (error) {
      console.error("Error al eliminar:", error);
    }
  };

  const handleEditar = (actividad) => {
    setEditando(actividad.id);
    setFormData({
      nombre: actividad.nombre,
      descripcion: actividad.descripcion,
      profesor: actividad.profesor,
      horarios: actividad.horarios.map(h => ({
        dia: h.dia,
        horarioInicio: h.horarioInicio,
        horarioFinal: h.horarioFinal,
        cupo: h.cupo
      }))
    });
  };

  const handleCancelar = () => {
    setEditando(null);
    setFormData({ nombre: '', descripcion: '', profesor: '', horarios: [] });
  };

  const handleActualizar = async (id) => {
    try {
      const tokenCookie = document.cookie.split('; ').find(row => row.startsWith('token='));
      const token = tokenCookie ? tokenCookie.split('=')[1] : null;

      const res = await fetch(`${process.env.REACT_APP_API_ACTIVIDADES_URL}/actividad/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify({
          nombre: formData.nombre,
          descripcion: formData.descripcion,
          profesor: formData.profesor,
          horarios: formData.horarios.map(h => ({
            dia: h.dia,
            horarioInicio: h.horarioInicio,
            horarioFinal: h.horarioFinal,
            cupo: parseInt(h.cupo)
          }))
        })
      });

      if (res.ok) {
        alert("Actividad actualizada correctamente");
        setEditando(null);
        fetchActividades();
      } else {
        alert("Error al actualizar la actividad");
      }
    } catch (error) {
      console.error("Error al actualizar:", error);
    }
  };

  return (
    <div className="w-full">
      {filtro && actividades && actividades.length === 0 ? (
        <Card className="border-destructive/50">
          <CardContent className="pt-6">
            <p className="text-center text-muted-foreground">
              ⚠ No se encontró ninguna actividad con el nombre buscado.
            </p>
          </CardContent>
        </Card>
      ) : (
        <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
          {actividades && actividades.map((act) => (
            <Card 
              key={act.id} 
              className="group hover:shadow-lg transition-all cursor-pointer border-border overflow-hidden"
              onClick={() => !editando && navigate(`/Detalle/${act.id}`)}
            >
              {editando === act.id ? (
                <div onClick={(e) => e.stopPropagation()}>
                  <CardHeader>
                    <Input
                      value={formData.nombre}
                      onChange={(e) => setFormData({ ...formData, nombre: e.target.value })}
                      placeholder="Nombre"
                      className="font-bold text-lg"
                    />
                  </CardHeader>
                  <CardContent className="space-y-4">
                    <textarea
                      value={formData.descripcion}
                      onChange={(e) => setFormData({ ...formData, descripcion: e.target.value })}
                      placeholder="Descripción"
                      className="w-full min-h-[80px] rounded-md border border-input bg-background px-3 py-2 text-sm resize-none"
                    />
                    <Input
                      value={formData.profesor}
                      onChange={(e) => setFormData({ ...formData, profesor: e.target.value })}
                      placeholder="Profesor"
                    />
                    
                    <div className="space-y-2">
                      <h4 className="font-semibold text-sm">Horarios:</h4>
                      {formData.horarios.map((h, i) => (
                        <div key={i} className="grid grid-cols-2 gap-2 p-2 border rounded-md">
                          <Input
                            value={h.dia}
                            onChange={(e) => {
                              const nuevosHorarios = [...formData.horarios];
                              nuevosHorarios[i].dia = e.target.value;
                              setFormData({ ...formData, horarios: nuevosHorarios });
                            }}
                            placeholder="Día"
                            className="col-span-2"
                          />
                          <Input
                            type="time"
                            value={h.horarioInicio}
                            onChange={(e) => {
                              const nuevosHorarios = [...formData.horarios];
                              nuevosHorarios[i].horarioInicio = e.target.value;
                              setFormData({ ...formData, horarios: nuevosHorarios });
                            }}
                          />
                          <Input
                            type="time"
                            value={h.horarioFinal}
                            onChange={(e) => {
                              const nuevosHorarios = [...formData.horarios];
                              nuevosHorarios[i].horarioFinal = e.target.value;
                              setFormData({ ...formData, horarios: nuevosHorarios });
                            }}
                          />
                          <Input
                            type="number"
                            value={h.cupo}
                            onChange={(e) => {
                              const nuevosHorarios = [...formData.horarios];
                              nuevosHorarios[i].cupo = e.target.value;
                              setFormData({ ...formData, horarios: nuevosHorarios });
                            }}
                            placeholder="Cupo"
                            className="col-span-2"
                          />
                        </div>
                      ))}
                    </div>
                  </CardContent>
                  <CardFooter className="flex gap-2">
                    <Button 
                      size="sm"
                      onClick={(e) => { e.stopPropagation(); handleActualizar(act.id); }}
                      className="flex-1"
                    >
                      <Save className="h-4 w-4 mr-2" />
                      Guardar
                    </Button>
                    <Button 
                      size="sm"
                      variant="outline"
                      onClick={(e) => { e.stopPropagation(); handleCancelar(); }}
                      className="flex-1"
                    >
                      <X className="h-4 w-4 mr-2" />
                      Cancelar
                    </Button>
                  </CardFooter>
                </div>
              ) : (
                <>
                  <CardHeader>
                    <div className="flex items-start justify-between">
                      <CardTitle className="text-xl font-mono uppercase tracking-tight">
                        {act.nombre}
                      </CardTitle>
                      <Badge variant="secondary" className="ml-2">
                        {act.horarios?.length || 0} horarios
                      </Badge>
                    </div>
                    <CardDescription className="flex items-center gap-2 mt-2">
                      <User className="h-4 w-4" />
                      {act.profesor}
                    </CardDescription>
                  </CardHeader>
                  <CardContent>
                    <p className="text-sm text-muted-foreground mb-4 line-clamp-2">
                      {act.descripcion}
                    </p>
                    
                    {act.horarios && act.horarios.length > 0 && (
                      <div className="space-y-2">
                        <div className="flex items-center gap-2 text-xs font-semibold text-muted-foreground">
                          <Clock className="h-3 w-3" />
                          Horarios disponibles
                        </div>
                        {act.horarios.slice(0, 2).map((h, idx) => (
                          <div key={idx} className="flex items-center justify-between text-sm p-2 rounded-md bg-secondary/50">
                            <div className="flex items-center gap-2">
                              <Calendar className="h-3 w-3" />
                              <span className="font-medium">{h.dia}</span>
                            </div>
                            <div className="flex items-center gap-2 text-xs text-muted-foreground">
                              <span>{h.horarioInicio} - {h.horarioFinal}</span>
                              <Badge variant="outline" className="ml-2">
                                <Users className="h-3 w-3 mr-1" />
                                {h.cupo}
                              </Badge>
                            </div>
                          </div>
                        ))}
                        {act.horarios.length > 2 && (
                          <p className="text-xs text-muted-foreground text-center">
                            +{act.horarios.length - 2} horarios más
                          </p>
                        )}
                      </div>
                    )}
                  </CardContent>
                  
                  {esAdmin && (
                    <CardFooter className="flex gap-2 bg-secondary/30" onClick={(e) => e.stopPropagation()}>
                      <Button
                        size="sm"
                        variant="outline"
                        onClick={(e) => {
                          e.stopPropagation();
                          handleEditar(act);
                        }}
                        className="flex-1"
                      >
                        <Edit className="h-4 w-4 mr-2" />
                        Editar
                      </Button>
                      <Button
                        size="sm"
                        variant="destructive"
                        onClick={(e) => {
                          e.stopPropagation();
                          handleEliminar(act.id);
                        }}
                        className="flex-1"
                      >
                        <Trash2 className="h-4 w-4 mr-2" />
                        Eliminar
                      </Button>
                    </CardFooter>
                  )}
                </>
              )}
            </Card>
          ))}
        </div>
      )}
    </div>
  );
}

export default ListadoActividades;
