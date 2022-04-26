var builder = WebApplication.CreateBuilder(args);

builder.Services.AddEndpointsApiExplorer();

var app = builder.Build();

app.UseHttpsRedirection();


app.MapGet("/", GetHome);

async Task<IResult> GetHome() {
    await Task.Delay(100);
    return Results.Ok("Ok");
}

app.Run();
