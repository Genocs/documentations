using System.Threading.Tasks;
using Dapr.Client;
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
        private readonly DaprClient daprClient;


        const string storeName = "statestore";
        const string key = "dotnetapp";
        const string externalService = "pythonapp";


        public OrderController(ILogger<HomeController> logger, DaprClient daprClient)
        {
            if (daprClient is null)
            {
                throw new System.ArgumentNullException(nameof(daprClient));
            }

            _logger = logger ?? throw new System.ArgumentNullException(nameof(logger));
            this.daprClient = daprClient;
        }

        [HttpPost("submit")]
        public async Task<IActionResult> PostSubmitOrderAsync([FromBody] OrderPayload payload)
        {
            if (payload == null || payload.Data == null)
            {
                return BadRequest("Empty payload");
            }

            await UpdateStateStorageAsync(payload);
            await CallExernalServiceAsync(payload);

            OrderConfirmationPayload response = new();
            response.Data = new OrderConfirmation { Id = payload.Data.Id, Description = payload.Data.Description, Status = "Acquired" };
            return Ok(response);
        }

        private async Task UpdateStateStorageAsync(OrderPayload payload)
        {
            await this.daprClient.SaveStateAsync(storeName, key, payload.Data);
            Order orderResult = await this.daprClient.GetStateAsync<Order>(storeName, key);
            _logger.LogInformation($"Order Payload: '{orderResult}'");
        }

        private async Task CallExernalServiceAsync(OrderPayload payload)
        {
            OrderConfirmationPayload externalServiceResult = await this.daprClient.InvokeMethodAsync<OrderPayload, OrderConfirmationPayload>(externalService, "order/submit", payload);
            _logger.LogInformation($"External Service Order Payload: '{externalServiceResult}'");
        }
    }
}
