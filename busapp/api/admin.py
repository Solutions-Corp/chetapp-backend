from django.contrib import admin
from .models import Ruta, Parada, Buseta, Favorito, Notificacion

admin.site.register(Ruta)
admin.site.register(Parada)
admin.site.register(Buseta)
admin.site.register(Favorito)
admin.site.register(Notificacion)
