# go-face-api

Permite actualizar tu estado de facebook y obtener informacion de tu perfil publico como nombre, genero, fecha de cumpleaños y una lista de las personas que son familiares contigo (hermanos/as, primos/as, madre, padre)

# Rutas

Dirigiendote a /permiso autorizas a la api a publicar en tu muro y obtener informacion de tu perfil publico y obtienes un token de acceso valido para dichos scopes. Luego si te diriges a /publicar debes suministrar el access token obtenido y el nuevo estado que deseas publicar
de esta forma:

{
"message":"tu nuevo estado",
"access_token":"el token valido"
}

Y por ultimo si nos dirigimos a /perfil debemos suministrar simplemente el access token obtenido para que obtenga la informacion de nuestro perfil publico, es decir de esta forma:

{
"access_token":"el token valido"
}
