using Genocs.Microservice.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace Genocs.Microservice.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class OrderController : ControllerBase
    {
        private readonly ILogger<HomeController> _logger;

        public OrderController(ILogger<HomeController> logger)
            => _logger = logger;

        [HttpPost("submit")]
        public IActionResult PostMyData([FromBody] Order order)
        {
            return Ok(new OrderConfirmation { Id = order.Id, Description = order.Description, Status = "Acquired" });
        }
    }

}
