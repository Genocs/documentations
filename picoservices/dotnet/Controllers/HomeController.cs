﻿using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace Genocs.Microservice.Controllers
{
    [ApiController]
    [Route("")]
    public class HomeController : ControllerBase
    {
        private readonly ILogger<HomeController> _logger;

        public HomeController(ILogger<HomeController> logger)
            => _logger = logger ?? throw new System.ArgumentNullException(nameof(logger));

        [HttpGet]
        public IActionResult Get()
            => Ok("Welcome to .NET!!!");

        [HttpGet("hc")]
        public IActionResult GetHealthCheck()
            => Ok("healthy");
    }
}
