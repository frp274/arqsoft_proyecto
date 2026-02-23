import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "../components/ui/card";
import { Button } from "../components/ui/button";
import { Input } from "../components/ui/input";
import { Label } from "../components/ui/label";
import { Badge } from "../components/ui/badge";
import { Dumbbell, ArrowRight, UserPlus, ShieldCheck, Mail, User } from "lucide-react";

function Registro() {
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    nombre: "",
    apellido: "",
    password: "",
    confirmPassword: ""
  });
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");

    if (
      !formData.username ||
      !formData.email ||
      !formData.nombre ||
      !formData.apellido ||
      !formData.password ||
      !formData.confirmPassword
    ) {
      setError("Todos los campos son obligatorios.");
      return;
    }

    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(formData.email)) {
      setError("El email no es válido.");
      return;
    }

    if (formData.password !== formData.confirmPassword) {
      setError("Las contraseñas no coinciden.");
      return;
    }

    if (formData.password.length < 6) {
      setError("La contraseña debe tener al menos 6 caracteres.");
      return;
    }

    setLoading(true);

    try {
      const response = await fetch(
        `${process.env.REACT_APP_API_USUARIOS_URL}/usuario`,
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            username: formData.username,
            email: formData.email,
            nombre: formData.nombre,
            apellido: formData.apellido,
            password: formData.password,
            es_admin: false
          })
        }
      );

      if (response.ok) {
        alert("¡Registro exitoso! Ahora puedes iniciar sesión.");
        navigate("/Login");
      } else {
        const errorData = await response.json().catch(() => ({}));
        setError(
          errorData.error ||
          errorData.message ||
          "Error al registrar el usuario. El nombre de usuario o email ya existe."
        );
      }
    } catch (error) {
      console.error("Error de registro:", error);
      setError("Error al conectar con el servidor.");
    } finally {
      setLoading(false);
    }
  };

  const irALogin = () => navigate("/Login");

  return (
    <div className="min-h-screen bg-gradient-to-br from-background via-secondary/20 to-primary/10 flex items-center justify-center p-4 py-12">
      <div className="w-full max-w-6xl grid md:grid-cols-2 gap-12 items-center">

        {/* Panel izquierdo: branding / info */}
        <div className="hidden md:block space-y-8">
          <div className="flex items-center gap-3">
            <div className="h-12 w-12 rounded-full bg-primary flex items-center justify-center shadow-[0_0_20px_-5px_hsl(var(--primary))]">
              <Dumbbell className="h-6 w-6 text-primary-foreground" />
            </div>
            <div>
              <h1 className="font-mono text-2xl font-bold uppercase tracking-tighter">
                GOOD GYM
              </h1>
              <Badge variant="secondary" className="mt-1">
                <UserPlus className="h-3 w-3 mr-1" />
                Join the Community
              </Badge>
            </div>
          </div>

          <div className="space-y-4">
            <h2 className="font-mono text-4xl font-bold uppercase tracking-tighter leading-tight">
              Start Your<br />Transformation Today
            </h2>
            <p className="text-lg text-muted-foreground max-w-md leading-relaxed">
              Crea tu perfil para acceder a las mejores clases, seguir tu progreso y
              formar parte de la comunidad fitness con más estilo.
            </p>
          </div>

          <div className="grid gap-4">
            {[
              { icon: ShieldCheck, text: "Acceso seguro y personalizado" },
              { icon: Mail, text: "Notificaciones de tus clases" },
              { icon: User, text: "Gestión fácil de agenda" }
            ].map((item, i) => (
              <div key={i} className="flex items-center gap-4 text-muted-foreground bg-secondary/30 p-3 rounded-xl border border-border/50">
                <item.icon className="h-5 w-5 text-primary" />
                <span className="font-medium text-white/80">{item.text}</span>
              </div>
            ))}
          </div>

          <div className="pt-4">
            <p className="text-sm text-muted-foreground">
              ¿Ya tenés una cuenta?{" "}
              <Button
                variant="link"
                className="p-0 h-auto font-semibold text-primary"
                onClick={irALogin}
              >
                Iniciá sesión aquí
                <ArrowRight className="ml-1 h-4 w-4" />
              </Button>
            </p>
          </div>
        </div>

        {/* Panel derecho: formulario */}
        <Card className="w-full shadow-[0_0_50px_-12px_hsl(var(--primary)/0.15)] border-primary/20 bg-card/60 backdrop-blur-xl hover:border-primary/40 transition-all duration-500">
          <CardHeader className="pb-4">
            <CardTitle className="font-mono text-3xl uppercase tracking-tighter text-transparent bg-clip-text bg-gradient-to-r from-white to-white/60">
              Crear cuenta
            </CardTitle>
            <CardDescription className="text-muted-foreground/80">
              Completá tus datos para empezar a entrenar
            </CardDescription>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleSubmit} className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2 group">
                  <Label htmlFor="nombre" className="text-xs text-muted-foreground group-focus-within:text-primary transition-colors">Nombre</Label>
                  <Input
                    id="nombre"
                    name="nombre"
                    placeholder="Tu nombre"
                    value={formData.nombre}
                    onChange={handleChange}
                    disabled={loading}
                    className="h-10 bg-background/50 border-input focus:border-primary/50 transition-all"
                  />
                </div>
                <div className="space-y-2 group">
                  <Label htmlFor="apellido" className="text-xs text-muted-foreground group-focus-within:text-primary transition-colors">Apellido</Label>
                  <Input
                    id="apellido"
                    name="apellido"
                    placeholder="Tu apellido"
                    value={formData.apellido}
                    onChange={handleChange}
                    disabled={loading}
                    className="h-10 bg-background/50 border-input focus:border-primary/50 transition-all"
                  />
                </div>
              </div>

              <div className="space-y-2 group">
                <Label htmlFor="email" className="text-xs text-muted-foreground group-focus-within:text-primary transition-colors">Email</Label>
                <Input
                  id="email"
                  name="email"
                  type="email"
                  placeholder="ejemplo@correo.com"
                  value={formData.email}
                  onChange={handleChange}
                  disabled={loading}
                  className="h-10 bg-background/50 border-input focus:border-primary/50 transition-all"
                />
              </div>

              <div className="space-y-2 group">
                <Label htmlFor="username" className="text-xs text-muted-foreground group-focus-within:text-primary transition-colors">Nombre de Usuario</Label>
                <Input
                  id="username"
                  name="username"
                  placeholder="Elige un alias"
                  value={formData.username}
                  onChange={handleChange}
                  disabled={loading}
                  className="h-10 bg-background/50 border-input focus:border-primary/50 transition-all"
                />
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2 group">
                  <Label htmlFor="password" className="text-xs text-muted-foreground group-focus-within:text-primary transition-colors">Contraseña</Label>
                  <Input
                    id="password"
                    name="password"
                    type="password"
                    placeholder="••••••"
                    value={formData.password}
                    onChange={handleChange}
                    disabled={loading}
                    className="h-10 bg-background/50 border-input focus:border-primary/50 transition-all"
                  />
                </div>
                <div className="space-y-2 group">
                  <Label htmlFor="confirmPassword" className="text-xs text-muted-foreground group-focus-within:text-primary transition-colors">Confirmar</Label>
                  <Input
                    id="confirmPassword"
                    name="confirmPassword"
                    type="password"
                    placeholder="••••••"
                    value={formData.confirmPassword}
                    onChange={handleChange}
                    disabled={loading}
                    className="h-10 bg-background/50 border-input focus:border-primary/50 transition-all"
                  />
                </div>
              </div>

              {error && (
                <div className="text-xs font-medium text-destructive bg-destructive/10 border border-destructive/30 rounded-md p-3 flex items-center gap-2 animate-in fade-in zoom-in duration-300">
                  <div className="w-1 h-1 rounded-full bg-destructive animate-pulse" />
                  {error}
                </div>
              )}

              <Button type="submit" disabled={loading} className="w-full h-12 font-bold text-lg uppercase tracking-wider relative overflow-hidden group shadow-[0_0_20px_-5px_hsl(var(--primary)/0.4)] hover:shadow-[0_0_30px_-5px_hsl(var(--primary)/0.6)] transition-all duration-300 mt-4">
                <span className="relative z-10">{loading ? "Registrando..." : "Crear Cuenta"}</span>
                <div className="absolute inset-0 bg-white/20 translate-y-full group-hover:translate-y-0 transition-transform duration-300 ease-in-out" />
              </Button>

              <div className="md:hidden text-center pt-4 border-t border-border/50">
                <p className="text-sm text-muted-foreground">
                  ¿Ya tenés una cuenta?{" "}
                  <Button
                    variant="link"
                    className="p-0 h-auto text-primary"
                    onClick={irALogin}
                    type="button"
                  >
                    Iniciá sesión aquí
                  </Button>
                </p>
              </div>
            </form>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}

export default Registro;
