from django.db import models
from django.contrib.auth.models import User

class Ruta(models.Model):
    nombre = models.CharField(max_length=100)
    descripcion = models.TextField()
    duracion_estimada = models.IntegerField()  # en minutos

    def __str__(self):
        return self.nombre

class Parada(models.Model):
    ruta = models.ForeignKey(Ruta, related_name='paradas', on_delete=models.CASCADE)
    nombre = models.CharField(max_length=100)
    latitud = models.DecimalField(max_digits=9, decimal_places=6)
    longitud = models.DecimalField(max_digits=9, decimal_places=6)
    orden = models.IntegerField()

    def __str__(self):
        return f"{self.nombre} ({self.ruta.nombre})"

class Buseta(models.Model):
    codigo = models.CharField(max_length=50)
    ruta = models.ForeignKey(Ruta, related_name='busetas', on_delete=models.CASCADE)
    latitud_actual = models.DecimalField(max_digits=9, decimal_places=6, null=True, blank=True)
    longitud_actual = models.DecimalField(max_digits=9, decimal_places=6, null=True, blank=True)
    estado = models.CharField(max_length=50, choices=[
        ('en_ruta', 'En Ruta'),
        ('detenido', 'Detenido'),
        ('fuera_de_servicio', 'Fuera de Servicio'),
    ], default='en_ruta')

    def __str__(self):
        return self.codigo

class Favorito(models.Model):
    usuario = models.ForeignKey(User, related_name='favoritos', on_delete=models.CASCADE)
    ruta = models.ForeignKey(Ruta, related_name='favoritos', on_delete=models.CASCADE)

    def __str__(self):
        return f"{self.usuario.username} - {self.ruta.nombre}"

class Notificacion(models.Model):
    usuario = models.ForeignKey(User, related_name='notificaciones', on_delete=models.CASCADE)
    mensaje = models.TextField()
    timestamp = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return f"Notificación para {self.usuario.username} en {self.timestamp}"
